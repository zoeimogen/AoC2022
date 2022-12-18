package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"strconv"
)

type cpuState struct {
	x          int
	cycles     int
	crt        string
}

func (s *cpuState) cycle() {
	if s.x - 1 <= (s.cycles % 40) && s.x + 1 >= (s.cycles % 40) {
		s.crt += "#"
	} else {
		s.crt += "."
	}
	if (s.cycles % 40) == 39 {
		s.crt += "\n"
	}
	s.cycles++
}

func (s *cpuState) execute(instruction string) {
    i := strings.Split(instruction, " ")
	if i[0] == "noop" {
		s.cycle()
	} else if i[0] == "addx" {
		s.cycle()
		s.cycle()
		arg, _ := strconv.Atoi(i[1])
        s.x += arg
	}
}


func runDay10(inputFile string) (int, string) {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer file.Close()

	scan := bufio.NewScanner(file)
	
	var part1 int
	var cpu cpuState
	cpu.x = 1
    breakpoint := 20

	for scan.Scan() {
		// Strong supicion we'll see this instruction set again, so it's an object.
		oldx := cpu.x
        cpu.execute(scan.Text())
		if cpu.cycles >= breakpoint {
			part1 += oldx * breakpoint
			breakpoint += 40
		}
	}

	return part1, cpu.crt
}

func main() {
	var inputFile = flag.String("input", "day10-input.txt", "Problem input file")
	flag.Parse()
	part1, part2 := runDay10(*inputFile)
	fmt.Printf("%d\n%s", part1, part2)
}
