package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

type loc [2]int

func considerMove(this loc, previous loc) (loc) {
	if this[0] < (previous[0] - 1) ||
	   this[0] > (previous[0] + 1) ||
	   this[1] < (previous[1] - 1) ||
	   this[1] > (previous[1] + 1) {
		if this[0] == previous[0] {
			// Yes, left/right move needed
			if this[1] > previous[1] { this[1]-- } else {this[1]++ }
		} else if this[1] == previous[1] {
			// Yes, up/down move needed
			if this[0] > previous[0] { this[0]-- } else {this[0]++ }
		} else {
			// Diagonal move needed.
			if this[0] > previous[0] { this[0]-- } else {this[0]++ }
			if this[1] > previous[1] { this[1]-- } else {this[1]++ }
		}
	}

	return this
}

//////////////////////////
// Part 1 object/functions
// This could be merged into part 2 if that was generalised to any number of tails.
type part1Data struct {
	head loc
	tail loc
	seen []loc
}

func (l *part1Data) seenLocation(a loc) bool {
    for _, b := range l.seen {
        if b == a {
            return true
        }
    }
    return false
}

func (l *part1Data) doPart1(dir byte, count int) {
    for i := 0; i < count; i++ {
		if dir == 'U' {
			l.head[0]--
		} else if dir == 'D' {
			l.head[0]++
		} else if dir == 'R' {
			l.head[1]++
		} else if dir == 'L' {
			l.head[1]--
		}
		l.tail = considerMove(l.tail, l.head)
		if !l.seenLocation(l.tail) {
			l.seen = append(l.seen, l.tail)
		} 
	}
}

//////////////////////////
// Part 2 object/functions
type part2Data struct {
	head loc
	tails [9]loc
	seen []loc
}

func (l *part2Data) seenLocation(a loc) bool {
    for _, b := range l.seen {
        if b == a {
            return true
        }
    }
    return false
}


func (l *part2Data) doPart2(dir byte, count int) {
    for i := 0; i < count; i++ {
		if dir == 'U' {
			l.head[0]--
		} else if dir == 'D' {
			l.head[0]++
		} else if dir == 'R' {
			l.head[1]++
		} else if dir == 'L' {
			l.head[1]--
		}
		for t := 0; t < len(l.tails); t++ {
			if t == 0 {
				l.tails[t] = considerMove(l.tails[t], l.head)
			} else {
			    l.tails[t] = considerMove(l.tails[t], l.tails[t-1])
			}
		}
		if !l.seenLocation(l.tails[len(l.tails)-1]) {
			l.seen = append(l.seen, l.tails[len(l.tails)-1])
		}
	}
}

func runDay09(inputFile string) (int, int) {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer file.Close()

	var part1 part1Data
	part1.seen = (append(part1.seen, part1.tail))
	var part2 part2Data
	part2.seen = (append(part2.seen, part2.tails[0]))

	scan := bufio.NewScanner(file)
	for scan.Scan() {
		var dir byte
		var count int
		n, _ := fmt.Sscanf(scan.Text(), "%c %d", &dir, &count)
		if n != 2 {
			fmt.Sprintf("Unable to parse: %s\n", scan.Text())
			os.Exit(1)
		}
        part1.doPart1(dir, count)
		part2.doPart2(dir, count)
	}

	return len(part1.seen), len(part2.seen)
}

func main() {
	var inputFile = flag.String("input", "day09-input.txt", "Problem input file")
	flag.Parse()
	part1, part2 := runDay09(*inputFile)
	fmt.Printf("%d %d\n", part1, part2)
}
