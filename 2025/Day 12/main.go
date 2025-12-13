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

func getNumRegions(s []string) int {
	numPossibleRegions := 0
	var shapeSizes []int
	shapeIdx := -1
	for _, line := range s {
		if len(line) > 1 {
			if line[0] >= '0' && line[0] <= '9' && line[1] == ':' {
				shapeIdx++
				shapeSizes = append(shapeSizes, 0)
			} else if line[0] == '#' || line[0] == '.' {
				shapeSizes[shapeIdx] += strings.Count(line, "#")
			} else {
				parts := strings.Split(line, ": ")
				width, _ := strconv.Atoi(strings.Split(parts[0], "x")[0])
				length, _ := strconv.Atoi(strings.Split(parts[0], "x")[1])
				nums := strings.Split(parts[1], " ")
				totalArea := 0
				for idx, num := range nums {
					intNum, _ := strconv.Atoi(num)
					totalArea += intNum * shapeSizes[idx]
				}
				if totalArea <= length*width {
					numPossibleRegions++
				}
			}
		}
	}
	return numPossibleRegions
}

func main() {
	filename := "tests.txt"
	output, err := ReadFile(filename)
	if err != nil {
		println(err)
	}

	fmt.Println(getNumRegions(output))
}
