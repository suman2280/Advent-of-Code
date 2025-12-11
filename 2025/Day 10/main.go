package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/draffensperger/golp"
)

func ReadFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var text []string

	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	return text, nil
}

func getFinalState(lights string) []int {
	finalState := make([]int, len(lights)-2)
	for i := 1; i < len(lights)-1; i++ {
		if lights[i] == '.' {
			finalState[i-1] = 0
		} else {
			finalState[i-1] = 1
		}
	}
	return finalState
}

func getButtons(buttons []string) [][]int {
	intButtons := make([][]int, len(buttons))
	for i, strButton := range buttons {
		values := strings.Split(strButton[1:len(strButton)-1], ",")
		intButtons[i] = make([]int, len(values))
		for j, value := range values {
			intButtons[i][j], _ = strconv.Atoi(value)
		}
	}
	return intButtons
}

func cmpSlice(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func getMinCombinations(lights string, buttons []string, _ string) int {
	finalState := getFinalState(lights)
	initialState := make([]int, len(finalState)) // by default all zeroes
	intButtons := getButtons(buttons)
	type Tuple struct {
		state  []int
		clicks int
	}
	visited := make(map[string]int)
	queue := make([]Tuple, 0)
	queue = append(queue, Tuple{
		state:  initialState,
		clicks: 0,
	})
	for len(queue) > 0 {
		tile := queue[0]  // Get first
		queue = queue[1:] // Remove it
		if cmpSlice(tile.state, finalState) {
			return tile.clicks
		}
		if _, ok := visited[fmt.Sprint(tile.state)]; ok {
			continue
		}
		visited[fmt.Sprint(tile.state)] = tile.clicks
		for _, button := range intButtons {
			newState := make([]int, len(tile.state))
			copy(newState, tile.state)
			for i := 0; i < len(button); i++ {
				if newState[button[i]] == 0 {
					newState[button[i]] = 1
				} else {
					newState[button[i]] = 0
				}
				if _, ok := visited[fmt.Sprint(newState)]; ok {
					continue
				}
				queue = append(queue, Tuple{
					state:  newState,
					clicks: tile.clicks + 1,
				})
			}
		}
	}
	return -1
}

func part1(s []string) int {
	sum := 0
	for _, line := range s {
		parts := strings.Split(line, " ")
		lights := parts[0]
		joltage := parts[len(parts)-1]
		buttons := parts[1 : len(parts)-1]
		sum += getMinCombinations(lights, buttons, joltage)
	}
	return sum
}

//  -- Part 2

func getFinalJoltage(joltages string) []int {
	strJoltages := strings.Split(joltages[1:len(joltages)-1], ",")
	finalJoltage := make([]int, len(strJoltages))
	for i := 0; i < len(strJoltages); i++ {
		finalJoltage[i], _ = strconv.Atoi(strJoltages[i])
	}
	return finalJoltage
}

func getMinCombinations2(_ string, buttons []string, joltages string) int {
	finalJoltage := getFinalJoltage(joltages)
	intButtons := getButtons(buttons)
	const maxClicks = 1000 // adjust as needed
	numJoltages := len(finalJoltage)
	lp := golp.NewLP(0, len(intButtons))
	lp.SetVerboseLevel(golp.NEUTRAL)
	objectiveCoeffs := make([]float64, len(intButtons))
	for i := 0; i < len(intButtons); i++ {
		objectiveCoeffs[i] = 1.0
		lp.SetInt(i, true)
		lp.SetBounds(i, 0.0, float64(maxClicks))
	}
	lp.SetObjFn(objectiveCoeffs)
	for i := 0; i < numJoltages; i++ {
		var entries []golp.Entry
		for j, btn := range intButtons {
			if slices.Contains(btn, i) {
				entries = append(entries, golp.Entry{Col: j, Val: 1.0})
			}
		}
		targetValue := float64(finalJoltage[i])
		if err := lp.AddConstraintSparse(entries, golp.EQ, targetValue); err != nil {
			panic(err)
		}
	}
	status := lp.Solve()
	if status != golp.OPTIMAL {
		return 0
	}
	solution := lp.Variables()
	clicks := 0
	for _, val := range solution {
		clicks += int(val + 0.5)
	}
	return clicks
}

func part2(s []string) int {
	sum := 0
	for _, line := range s {
		parts := strings.Split(line, " ")
		lights := parts[0]
		joltage := parts[len(parts)-1]
		buttons := parts[1 : len(parts)-1]
		sum += getMinCombinations2(lights, buttons, joltage)
	}
	return sum
}

func main() {
	filename := "tests.txt"
	output, err := ReadFile(filename)
	if err != nil {
		println(err)
	}

	fmt.Println("Part1:", part1(output))
	fmt.Println("Part2:", part2(output))
}
