package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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

func replaceAt(s string, idx int, char rune) string {
	runes := []rune(s)
	if idx >= 0 && idx < len(runes) {
		runes[idx] = char
		return string(runes)
	}
	return s
}

func part1(s []string) int {
	total := 0
	for y := 1; y < len(s); y++ {
		for x := 0; x < len(s[y]); x++ {
			if s[y-1][x] == 'S' || s[y-1][x] == '|' {
				if s[y][x] == '^' {
					total++
					s[y] = replaceAt(s[y], x+1, '|')
					s[y] = replaceAt(s[y], x-1, '|')
				} else {
					s[y] = replaceAt(s[y], x, '|')
				}
			}
		}
	}
	return total
}

func recursive(s []string, x, y int, cache map[string]int) int {
	if y == len(s) {
		return 1
	} else {
		for ; y < len(s); y++ {
			if s[y][x] == '^' {
				if v, ok := cache[fmt.Sprintf("%d,%d", x, y)]; ok {
					return v
				} else {
					v = recursive(s, x-1, y, cache) + recursive(s, x+1, y, cache)
					cache[fmt.Sprintf("%d,%d", x, y)] = v
					return v
				}
			}
		}
	}
	return 1
}

func part2(s []string) int {
	total := 0
	y, x := 0, strings.Index(s[0], "S")
	cache := make(map[string]int)
	total += recursive(s, x, y, cache)
	return total
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
