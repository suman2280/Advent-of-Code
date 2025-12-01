package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func readFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var text []string

	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	file.Close()
	return text, nil
}

func getPassword1(direction []string) int {
	count := 0
	point := 50
	for _, el := range direction {
		switch el[0] {
		case 'L':
			numPart := el[1:]
			num, err := strconv.Atoi(numPart)
			if err != nil {
				fmt.Println(err)
			}
			point = (point - num) % 100
			if point == 0 {
				count++
			}
		case 'R':
			numPart := el[1:]
			num, err := strconv.Atoi(numPart)
			if err != nil {
				fmt.Println(err)
			}
			point = (point + num) % 100
			if point == 0 {
				count++
			}
		}
	}

	return count
}

func getPassword2(direction []string) int {
	count := 0
	point := 50
	for _, el := range direction {
		switch el[0] {
		case 'L':
			numPart := el[1:]
			num, err := strconv.Atoi(numPart)
			if err != nil {
				fmt.Println(err)
			}
			for range num {
				point--
				point = (point%100 + 100) % 100
				if point == 0 {
					count++
				}
			}
		case 'R':
			numPart := el[1:]
			num, err := strconv.Atoi(numPart)
			if err != nil {
				fmt.Println(err)
			}
			for range num {
				point++
				point = point % 100
				if point == 0 {
					count++
				}
			}
		}
	}

	return count
}

func main() {
	text, err := readFile("tests.txt")
	if err != nil {
		println(err)
	}
	fmt.Println("The password for part one is", getPassword1(text))
	fmt.Println("The password for part two is", getPassword2(text))
}
