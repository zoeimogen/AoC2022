package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func splitInput(data []byte, atEOF bool) (advance int, token []byte, err error) {
	for i := 0; i < len(data); i++ {
		if data[i] == '\n' && data[i-1] == '\n' {
			return i + 1, data[:i-1], nil
		}
	}
	if !atEOF {
		return 0, nil, nil
	}
	return 0, data, bufio.ErrFinalToken
}

func loadState(input string) []string {
	lines := strings.Split(input, "\n")
	slots := len(lines[0])/4 + 1

	var state = make([]string, slots)

	for _, line := range lines {
		for i := 0; i < slots; i++ {
			c := line[i*4+1]
			if c > byte('9') {
				state[i] = string(c) + state[i]
			}
		}
	}
	return state
}

func doPart1Move(state []string, move string) []string {
	var m int
	var f int
	var t int

	if move == "" {
		return state
	}

	n, _ := fmt.Sscanf(move, "move %d from %d to %d", &m, &f, &t)
	if n != 3 {
		fmt.Printf("Can't parse move: %s", move)
		return state
	}

	from := state[f-1]
	to := state[t-1]

	for i := 0; i < m; i++ {
		to += string(from[len(from)-i-1])
	}

	state[f-1] = from[:len(from)-m]
	state[t-1] = to

	return state
}

func doPart2Move(state []string, move string) []string {
	var m int
	var f int
	var t int

	if move == "" {
		return state
	}

	n, _ := fmt.Sscanf(move, "move %d from %d to %d", &m, &f, &t)
	if n != 3 {
		fmt.Printf("Can't parse move: %s", move)
		return state
	}

	from := state[f-1]
	to := state[t-1]

	state[f-1] = from[:len(from)-m]
	state[t-1] = to + from[len(from)-m:]

	return state
}

func runDay05(inputFile string) (string, string) {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer file.Close()

	scan := bufio.NewScanner(file)
	scan.Split(splitInput)
	scan.Scan()
	initial := scan.Text()
	part1State := loadState(initial)
	part2State := loadState(initial) // Yeah, this is lazy but it's a fast function

	scan.Scan()
	moves := strings.Split(scan.Text(), "\n")

	for _, m := range moves {
		part1State = doPart1Move(part1State, m)
		part2State = doPart2Move(part2State, m)
	}

	part1 := ""
	for _, l := range part1State {
		part1 += string(l[len(l)-1])
	}

	part2 := ""
	for _, l := range part2State {
		part2 += string(l[len(l)-1])
	}

	return part1, part2
}

func main() {
	var inputFile = flag.String("input", "day05-input.txt", "Problem input file")
	flag.Parse()
	part1, part2 := runDay05(*inputFile)
	fmt.Printf("%s, %s\n", part1, part2)
}
