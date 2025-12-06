package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

func part1(s []string) int {
	solutions := make([]int, len(s[len(s)-1]))
	operations := strings.Fields(s[len(s)-1])
	for i, line := range s[:len(s)-1] {
		numbers := strings.Fields(line)
		for j, number := range numbers {
			n, _ := strconv.Atoi(number)
			if i == 0 {
				solutions[j] = n
			} else if operations[j] == "+" {
				solutions[j] = solutions[j] + n
			} else {
				solutions[j] = solutions[j] * n
			}
		}
	}
	sum := 0
	for _, solution := range solutions {
		sum += solution
	}
	return sum
}

func part2(s []string) int {
	var (
		newColumn []int
		totalSum  int = 0
	)
	for x := len(s[0]) - 1; x >= 0; x-- {
		number := 0
		isOperated := false
		for y := range s {
			if s[y][x] == ' ' {
				continue
			} else if s[y][x] == '+' {
				result := number
				for i := 0; i < len(newColumn); i++ {
					result += newColumn[i]
				}
				totalSum += result
				isOperated = true
				newColumn = []int{}
				x--
				break
			} else if s[y][x] == '*' {
				result := number
				for i := 0; i < len(newColumn); i++ {
					result *= newColumn[i]
				}
				totalSum += result
				isOperated = true
				newColumn = []int{}
				x--
				break
			} else {
				n, _ := strconv.Atoi(string(s[y][x]))
				number = number*10 + n
			}
		}
		if !isOperated {
			newColumn = append(newColumn, number)
		}
	}
	return totalSum
}

func main() {
	filename := "tests.txt"
	output, err := ReadFile(filename)
	if err != nil {
		println(err)
	}

	fmt.Println(part1(output))
	fmt.Println(part2(output))
}
