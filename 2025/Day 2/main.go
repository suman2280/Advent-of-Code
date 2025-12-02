package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func part1(productIds []string) int {
	sum := 0

	for _, p := range productIds {
		idParts := strings.Split(p, "-")
		firstId, _ := strconv.Atoi(idParts[0])
		lastId, _ := strconv.Atoi(idParts[1])

		for id := firstId; id <= lastId; id++ {
			strId := strconv.Itoa(id)
			if len(strId)%2 != 0 {
				continue
			}

			mid := len(strId) / 2
			if strId[:mid] == strId[mid:] {
				sum += id
			}
		}
	}

	return sum
}

func part2(productIds []string) int {
	sum := 0

	for _, p := range productIds {
		idParts := strings.Split(p, "-")
		firstId, _ := strconv.Atoi(idParts[0])
		lastId, _ := strconv.Atoi(idParts[1])

		for id := firstId; id <= lastId; id++ {
			strId := strconv.Itoa(id)
			if hasRepeatingPattern(strId) {
				sum += id
			}
		}
	}
	return sum
}

func hasRepeatingPattern(str string) bool {
outer:
	for n := 1; n <= len(str)/2; n++ {
		if len(str)%n != 0 {
			continue
		}
		a := str[:n]
		for i := n; i <= len(str)-n; i += n {
			if a != str[i:i+n] {
				continue outer
			}
		}
		return true
	}
	return false
}

func main() {
	file, _ := os.Open("tests.txt")
	defer file.Close()
	sc := bufio.NewScanner(file)
	sc.Scan()
	productIds := strings.Split(sc.Text(), ",")

	fmt.Println("part1:", part1(productIds))
	fmt.Println("part2:", part2(productIds))
}
