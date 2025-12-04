package main

import (
	"bufio"
	"fmt"
	"os"
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

func part1(rolls []string) int {
	if len(rolls) == 0 {
		return 0
	}

	rows := len(rolls)
	cols := len(rolls[0])

	isValid := func(r, c int) bool {
		return r >= 0 && r < rows && c >= 0 && c < cols
	}

	count := 0

	for r := range rows {
		for c := range cols {
			if rolls[r][c] == '@' {
				neighbor := 0

				direction := []struct{ dr, dc int }{
					{-1, -1},
					{-1, 0},
					{-1, 1},
					{0, -1},
					{0, 1},
					{1, -1},
					{1, 0},
					{1, 1},
				}

				for _, dir := range direction {
					nr, nc := r+dir.dr, c+dir.dc

					if isValid(nr, nc) {
						if rolls[nr][nc] == '@' {
							neighbor++
						}
					}
				}
				if neighbor < 4 {
					count++
				}
			}
		}
	}
	return count
}

func part2(rolls []string) int {
	if len(rolls) == 0 {
		return 0
	}

	current := gridToRune(rolls)
	count := 0

	for {
		removed, nextGrid := removeRolls(current)

		if removed == 0 {
			break
		}

		count += removed
		current = nextGrid
	}
	return count
}

func gridToRune(grid []string) [][]rune {
	rows := len(grid)
	if rows == 0 {
		return nil
	}
	runeGrid := make([][]rune, rows)
	for r, row := range grid {
		runeGrid[r] = []rune(row)
	}
	return runeGrid
}

func removeRolls(grid [][]rune) (int, [][]rune) {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0, grid
	}

	rows := len(grid)
	cols := len(grid[0])
	rollsToRemove := make([][2]int, 0)

	newGrid := make([][]rune, rows)
	for r := range grid {
		newGrid[r] = make([]rune, cols)
		copy(newGrid[r], grid[r])
	}

	isValid := func(r, c int) bool {
		return r >= 0 && r < rows && c >= 0 && c < cols
	}

	for r := range rows {
		for c := range cols {
			if grid[r][c] == '@' {

				neighborRolls := 0
				directions := []struct{ dr, dc int }{
					{-1, -1},
					{-1, 0},
					{-1, 1},
					{0, -1},
					{0, 1},
					{1, -1},
					{1, 0},
					{1, 1},
				}

				for _, dir := range directions {
					nr, nc := r+dir.dr, c+dir.dc

					if isValid(nr, nc) {
						if grid[nr][nc] == '@' {
							neighborRolls++
						}
					}
				}

				if neighborRolls < 4 {
					rollsToRemove = append(rollsToRemove, [2]int{r, c})
				}
			}
		}
	}

	for _, pos := range rollsToRemove {
		r, c := pos[0], pos[1]
		newGrid[r][c] = '.'
	}

	return len(rollsToRemove), newGrid
}

func main() {
	rolls, err := ReadFile("tests.txt")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("part1:", part1(rolls))
	fmt.Println("part2:", part2(rolls))
}
