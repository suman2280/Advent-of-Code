package main

import (
	"bufio"
	"fmt"
	"math/big"
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

	file.Close()
	return text, nil
}

func part1(banks []string) int {
	total := 0

	for _, bank := range banks {
		maxVolt := 0
		var batteries []int
		for _, char := range bank {
			if char >= '0' && char <= '9' {
				batteries = append(batteries, int(char-'0'))
			}
		}
		for i := 0; i < len(batteries); i++ {
			first := batteries[i]
			for j := i + 1; j < len(batteries); j++ {
				second := batteries[j]
				current := first*10 + second

				if current > maxVolt {
					maxVolt = current
				}
			}
		}
		total += maxVolt
	}

	return total
}

func part2(banks []string) string {
	total := big.NewInt(0)
	const required = 12

	for _, bank := range banks {
		str := largestSequence(bank, required)
		val := big.NewInt(0)
		val.SetString(str, 10)
		total.Add(total, val)
	}
	return total.String()
}

func largestSequence(bank string, required int) string {
	remove := len(bank) - required

	if remove < 0 {
		return bank
	}

	var res []rune
	for _, val := range bank {
		for len(res) > 0 && remove > 0 && val > res[len(res)-1] {
			res = res[:len(res)-1]
			remove--
		}
		res = append(res, val)
	}

	if remove > 0 {
		res = res[:len(res)-remove]
	}
	if len(res) > required {
		res = res[:required]
	}

	return string(res)
}

func main() {
	banks, err := ReadFile("tests.txt")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("part1:", part1(banks))
	fmt.Println("part2:", part2(banks))
}
