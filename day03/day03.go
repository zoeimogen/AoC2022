package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

func getPriority(data rune) int {
	if data > 'Z' {
		return int(data - 'a' + 1)
	} else {
		return int(data - 'A' + 27)
	}
}

func calculateScorePart1(data string) int {
	l := len(data)
	first := data[0 : l/2]
	second := data[l/2 : l]
	for _, c := range first {
		for _, d := range second {
			if c == d {
				return getPriority(c)
			}
		}
	}

	fmt.Printf("Unable to find a matching token in compartments\n")
	return 0
}

func calculateScorePart2(data []string) int {
	var appearances = [52]int{}

	for _, d := range data {
		var thisPack = [52]int{}
		for _, c := range d {
			p := getPriority(c) - 1
			if thisPack[p] == 0 {
				appearances[getPriority(c)-1] += 1
				thisPack[getPriority(c)-1] = 1
			}
		}
	}

	for i, c := range appearances {
		if c == 3 {
			return i + 1
		}
	}

	fmt.Printf("Unable to find a matching token in group\n")
	return 0
}

func runDay03(inputFile string) (int, int) {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer file.Close()

	var part1 = 0
	var part2 = 0
	var group = []string{}
	scan := bufio.NewScanner(file)
	for scan.Scan() {
		part1 += calculateScorePart1(scan.Text())
		group = append(group, scan.Text())
		if len(group) == 3 {
			part2 += calculateScorePart2(group)
			group = []string{}
		}
	}

	return part1, part2
}

func main() {
	var inputFile = flag.String("input", "day03-input.txt", "Problem input file")
	flag.Parse()
	part1, part2 := runDay03(*inputFile)
	fmt.Printf("%d, %d\n", part1, part2)
}
