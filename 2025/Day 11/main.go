package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	You    = Device("you")
	Server = Device("svr")
	Fft    = Device("fft")
	Dac    = Device("dac")
	Out    = Device("out")
)

type Device string

type Rack map[Device][]Device

func (r Rack) countPathsToOut(start Device) int {
	memo := make(map[Device]int)
	var count func(Device) int
	count = func(st Device) int {
		if v, ok := memo[st]; ok {
			return v
		}
		if st == Out {
			memo[st] = 1
			return 1
		}

		nextDevices, _ := r[st]
		sum := 0
		for _, next := range nextDevices {
			sum += count(next)
		}
		memo[st] = sum
		return sum
	}

	return count(start)
}

type State struct {
	pos       Device
	passedFft bool
	passedDac bool
}

func (r Rack) countFftDacPathsToOut(start Device) int {
	memo := make(map[State]int)

	var count func(State) int
	count = func(state State) int {
		if v, ok := memo[state]; ok {
			return v
		}

		if state.pos == Out {
			if state.passedDac && state.passedFft {
				memo[state] = 1
				return 1
			} else {
				memo[state] = 0
				return 0
			}
		}

		if state.pos == Fft {
			state.passedFft = true
		}
		if state.pos == Dac {
			state.passedDac = true
		}

		nextDevices, _ := r[state.pos]
		sum := 0
		for _, next := range nextDevices {
			new := state
			new.pos = next
			sum += count(new)
		}
		memo[state] = sum
		return sum
	}

	return count(State{pos: start})
}

func parseInput(input string) Rack {
	trimmed := strings.TrimSpace(input)
	lines := strings.Split(trimmed, "\n")

	rack := make(Rack)
	for _, line := range lines {
		semic := strings.Index(line, ":")
		dev := Device(line[:semic])

		parts := strings.SplitSeq(line[semic+2:], " ")
		for part := range parts {
			rack[dev] = append(rack[dev], Device(part))
		}
	}
	return rack
}

func main() {
	input, err := os.ReadFile("tests.txt")
	if err != nil {
		fmt.Println(err)
	}
	rack := parseInput(string(input))

	fmt.Println("Part1:", rack.countPathsToOut(You))
	fmt.Println("Part2:", rack.countFftDacPathsToOut(Server))
}
