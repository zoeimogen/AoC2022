package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func getBlocks(data []byte, atEOF bool) (advance int, token []byte, err error) {
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

func totalBlock(data string) int {
	var total int = 0
	lines := strings.Split(data, "\n")
	for _, d := range lines {
		v, _ := strconv.Atoi(d)
		total += v
	}
	return total
}

func runDay01(inputFile string) (int, int) {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer file.Close()

	var elfList []int
	scan := bufio.NewScanner(file)
	scan.Split(getBlocks)
	for scan.Scan() {
		thisElf := totalBlock(scan.Text())
		elfList = append(elfList, thisElf)
	}

	sort.Ints(elfList[:])
	l := len(elfList)
	return elfList[l-1], elfList[l-1] + elfList[l-2] + elfList[l-3]
}

func main() {
	var inputFile = flag.String("input", "day01-input.txt", "Problem input file")
	flag.Parse()
	part1, part2 := runDay01(*inputFile)
	fmt.Printf("%d, %d\n", part1, part2)
}
