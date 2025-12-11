package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
)

type junctionBoxes struct {
	p1, p2, dist int
}

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

func part1(s []string) int {
	var x1, x2, y1, y2, z1, z2 int
	var jBoxes []junctionBoxes
	for i := 0; i < len(s)-1; i++ {
		_, _ = fmt.Sscanf(s[i], "%d,%d,%d", &x1, &y1, &z1)
		for j := i + 1; j < len(s); j++ {
			_, _ = fmt.Sscanf(s[j], "%d,%d,%d", &x2, &y2, &z2)
			distance := ((x2 - x1) * (x2 - x1)) + ((y2 - y1) * (y2 - y1)) + ((z2 - z1) * (z2 - z1))
			jBoxes = append(jBoxes, junctionBoxes{i, j, distance})
		}
	}
	sort.Slice(jBoxes, func(i, j int) bool {
		return jBoxes[i].dist < jBoxes[j].dist
	})
	if len(jBoxes) >= 1000 {
		jBoxes = jBoxes[:1000]
	} else {
		jBoxes = jBoxes[:10]
	}
	circuits := make([][]int, 0)
	var foundP1, foundP2 int
	for i := 0; i < len(jBoxes); i++ {
		foundP1, foundP2 = -1, -1
		for j := 0; j < len(circuits); j++ {
			if slices.Contains(circuits[j], jBoxes[i].p1) {
				foundP1 = j
			}
			if slices.Contains(circuits[j], jBoxes[i].p2) {
				foundP2 = j
			}
		}
		if foundP1 == -1 && foundP2 == -1 {
			circuits = append(circuits, []int{jBoxes[i].p1, jBoxes[i].p2})
		} else if foundP1 == foundP2 {
			continue
		} else if foundP1 != -1 && foundP2 == -1 {
			circuits[foundP1] = append(circuits[foundP1], jBoxes[i].p2)
		} else if foundP1 == -1 && foundP2 != -1 {
			circuits[foundP2] = append(circuits[foundP2], jBoxes[i].p1)
		} else if foundP1 != -1 && foundP2 != -1 && foundP1 != foundP2 {
			circuits[foundP1] = append(circuits[foundP1], circuits[foundP2]...) // add p2 on p1
			circuits = append(circuits[:foundP2], circuits[foundP2+1:]...)      // remove p2
		}
	}
	slices.SortFunc(circuits, func(c1, c2 []int) int {
		return len(c2) - len(c1)
	})
	return len(circuits[0]) * len(circuits[1]) * len(circuits[2])
}

func part2(s []string) int {
	var x1, x2, y1, y2, z1, z2 int
	var jBoxes []junctionBoxes
	for i := 0; i < len(s)-1; i++ {
		_, _ = fmt.Sscanf(s[i], "%d,%d,%d", &x1, &y1, &z1)
		for j := i + 1; j < len(s); j++ {
			_, _ = fmt.Sscanf(s[j], "%d,%d,%d", &x2, &y2, &z2)
			distance := ((x2 - x1) * (x2 - x1)) + ((y2 - y1) * (y2 - y1)) + ((z2 - z1) * (z2 - z1))
			jBoxes = append(jBoxes, junctionBoxes{i, j, distance})
		}
	}
	sort.Slice(jBoxes, func(i, j int) bool {
		return jBoxes[i].dist < jBoxes[j].dist
	})
	circuits := make([][]int, 0)
	var foundP1, foundP2 int
	result := 0
	for i := 0; i < len(jBoxes); i++ {
		foundP1, foundP2 = -1, -1
		for j := 0; j < len(circuits); j++ {
			if slices.Contains(circuits[j], jBoxes[i].p1) {
				foundP1 = j
			}
			if slices.Contains(circuits[j], jBoxes[i].p2) {
				foundP2 = j
			}
		}
		if foundP1 == -1 && foundP2 == -1 {
			circuits = append(circuits, []int{jBoxes[i].p1, jBoxes[i].p2})
		} else if foundP1 == foundP2 {
			continue
		} else if foundP1 != -1 && foundP2 == -1 {
			circuits[foundP1] = append(circuits[foundP1], jBoxes[i].p2)
		} else if foundP1 == -1 && foundP2 != -1 {
			circuits[foundP2] = append(circuits[foundP2], jBoxes[i].p1)
		} else if foundP1 != -1 && foundP2 != -1 && foundP1 != foundP2 {
			circuits[foundP1] = append(circuits[foundP1], circuits[foundP2]...) // add p2 on p1
			circuits = append(circuits[:foundP2], circuits[foundP2+1:]...)      // remove p2
		}
		_, _ = fmt.Sscanf(s[jBoxes[i].p1], "%d,%d,%d", &x1, &y1, &z1)
		_, _ = fmt.Sscanf(s[jBoxes[i].p2], "%d,%d,%d", &x2, &y2, &z2)
		result = x1 * x2
	}
	return result
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
