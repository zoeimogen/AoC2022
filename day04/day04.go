package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func runDay04(inputFile string) (int, int) {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer file.Close()

	scan := bufio.NewScanner(file)
	part1 := 0
	part2 := 0

	for scan.Scan() {
		d := strings.Split(scan.Text(), ",")
		a := strings.Split(d[0], "-")
		x := strings.Split(d[1], "-")
		b, _ := strconv.Atoi(a[0])
		c, _ := strconv.Atoi(a[1])
		y, _ := strconv.Atoi(x[0])
		z, _ := strconv.Atoi(x[1])

		if (b <= y && c >= z) || (b >= y && c <= z) {
			part1 += 1
			part2 += 1
		} else if (y <= b && b <= z) || (y <= c && c <= z) {
			part2 += 1
		}
	}

	return part1, part2
}

func main() {
	var inputFile = flag.String("input", "day04-input.txt", "Problem input file")
	flag.Parse()
	part1, part2 := runDay04(*inputFile)
	fmt.Printf("%d, %d\n", part1, part2)
}
