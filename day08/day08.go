package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

func look(trees [][]int, h int, r int, c int, dr int, dc int) bool {
	newR := r + dr
	newC := c + dc

	if newR < 0 || newC < 0 || newR >= len(trees) || newC >= len(trees[newR]) {
		// Tree is visible
		return true
	}

	if trees[newR][newC] >= h {
		// Tree is hidden in this direction
		return false
	}
	return look(trees, h, newR, newC, dr, dc)
}

func distance(trees [][]int, h int, r int, c int, dr int, dc int) int {
	newR := r + dr
	newC := c + dc

	if newR < 0 || newC < 0 || newR >= len(trees) || newC >= len(trees[newR]) {
		return 0
	}

	if trees[newR][newC] >= h {
		return 1
	}

	return distance(trees, h, newR, newC, dr, dc) + 1
}

func runDay08(inputFile string) (int, int) {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer file.Close()

	var trees [][]int

	scan := bufio.NewScanner(file)
	for scan.Scan() {
		var row []int
		for _, tree := range scan.Text() {
			t, _ := strconv.Atoi(string(tree))
			row = append(row, t)
		}
		trees = append(trees, row)
	}

	var part1 int
	var part2 int

	for r, row := range trees {
		for c, tree := range row {
			if look(trees, tree, r, c, -1, 0) || look(trees, tree, r, c, 1, 0) ||
				look(trees, tree, r, c, 0, -1) || look(trees, tree, r, c, 0, 1) {
				part1 += 1
			}
			scenic := distance(trees, tree, r, c, -1, 0)
			scenic *= distance(trees, tree, r, c, 1, 0)
			scenic *= distance(trees, tree, r, c, 0, -1)
			scenic *= distance(trees, tree, r, c, 0, 1)
			if scenic > part2 {
				part2 = scenic
			}
		}
	}
	return part1, part2
}

func main() {
	var inputFile = flag.String("input", "day08-input.txt", "Problem input file")
	flag.Parse()
	part1, part2 := runDay08(*inputFile)
	fmt.Printf("%d %d\n", part1, part2)
}
