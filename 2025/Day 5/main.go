package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func readFile(filename string) ([][]int, []int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	var ranges [][]int
	var availableId []int

	scanner := bufio.NewScanner(file)

	isRangeSection := true

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if line == "" {
			isRangeSection = false
			continue
		}

		if isRangeSection {
			parts := strings.Split(line, "-")

			startStr := strings.TrimSpace(parts[0])
			endStr := strings.TrimSpace(parts[1])

			start, _ := strconv.Atoi(startStr)
			end, _ := strconv.Atoi(endStr)

			ranges = append(ranges, []int{start, end})
		} else {
			idStr := strings.TrimSpace(line)
			id, _ := strconv.Atoi(idStr)
			availableId = append(availableId, id)
		}
	}
	return ranges, availableId, nil
}

func part1(ranges [][]int, availableId []int) int {
	fresh := 0

	for _, id := range availableId {
		isFresh := false
		for _, r := range ranges {
			if id >= r[0] && id <= r[1] {
				isFresh = true
				break
			}
		}
		if isFresh {
			fresh++
		}
	}

	return fresh
}

func part2(ranges [][]int) int {
	fresh := 0
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i][0] < ranges[j][0]
	})

	merged := [][]int{ranges[0]}

	for i := 1; i < len(ranges); i++ {
		lastMerged := merged[len(merged)-1]

		if ranges[i][0] <= lastMerged[1] {
			if ranges[i][1] > lastMerged[1] {
				merged[len(merged)-1][1] = ranges[i][1]
			}
		} else {
			merged = append(merged, ranges[i])
		}
	}

	for _, r := range merged {
		fresh += r[1] - r[0] + 1
	}

	return fresh
}

func main() {
	ranges, availableId, err := readFile("tests.txt")
	if err != nil {
		fmt.Println("err")
	}

	fmt.Println("Part1:", part1(ranges, availableId))
	fmt.Println("Part2:", part2(ranges))
}
