package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func calculateScorePart1(round string) int {
	// thing := []string{"Rock    ", "Paper   ", "Scissors"}
	d := strings.Split(round, " ")
	if len(d) != 2 {
		fmt.Printf("Unexpected input: %s\n", round)
		return 0
	}

	them := int([]byte(d[0])[0] - 'A')
	us := int([]byte(d[1])[0] - 'X')

	// Score for the play we made
	score := us + 1
	if (them == us-1) || (them == 2 && us == 0) { // We win
		score += 6
	} else if them == us { // Draw
		score += 3
	}

	// fmt.Printf("Round Score: %d (%s %s %d | %s %s %d)\n", score, d[0], thing[them], them, d[1], thing[us], us)
	return score
}

func calculateScorePart2(round string) int {
	d := strings.Split(round, " ")
	if len(d) != 2 {
		fmt.Printf("Unexpected input: %s\n", round)
		return 0
	}

	them := int([]byte(d[0])[0] - 'A')

	if d[1] == "X" { // We need to lose
		if them == 0 { // They played Rock, we play Scissors
			return 3
		} else {
			return them
		}
	} else if d[1] == "Y" { // We need to draw, so play the same as them
		return 4 + them
	} else { // We need to win
		if them == 2 { // They played Scissors, we play Rock
			return 7
		} else {
			return them + 8
		}
	}
}

func runDay02(inputFile string) (int, int) {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer file.Close()

	var part1 = 0
	var part2 = 0
	scan := bufio.NewScanner(file)
	for scan.Scan() {
		part1 += calculateScorePart1(scan.Text())
		part2 += calculateScorePart2(scan.Text())
	}

	return part1, part2
}

func main() {
	var inputFile = flag.String("input", "day02-input.txt", "Problem input file")
	flag.Parse()
	part1, part2 := runDay02(*inputFile)
	fmt.Printf("%d, %d\n", part1, part2)
}
