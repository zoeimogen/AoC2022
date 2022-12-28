package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func runPart1(terrain [][]byte, start [2]int, finish[2]int) (int) {
	// Yes, we could use channels here but it overcomplicates a simple problem
	queue := [][3]int{{start[0], start[1], 0}}
    visited := make([]bool, len(terrain)*len(terrain[0]))

	for len(queue) > 0 {
		var xyz [3]int
		xyz, queue = queue[0], queue[1:]
		if xyz[0] == finish[0] && xyz[1] == finish[1] {
			return xyz[2]
		}

		// Loop prevention
		index := xyz[0]*len(terrain[0]) + xyz[1]
		if visited[index] {
			continue
		}
		visited[index] = true

		thisTerrain := terrain[xyz[0]][xyz[1]]
		if xyz[0] > 0 && terrain[xyz[0]-1][xyz[1]] <= thisTerrain+1 {
			queue = append(queue, [3]int{xyz[0]-1, xyz[1], xyz[2]+1})
		}
		if xyz[0] < len(terrain)-1 && terrain[xyz[0]+1][xyz[1]] <= thisTerrain+1 {
			queue = append(queue, [3]int{xyz[0]+1, xyz[1], xyz[2]+1})
		}
		if xyz[1] > 0 && terrain[xyz[0]][xyz[1]-1] <= thisTerrain+1 {
			queue = append(queue, [3]int{xyz[0], xyz[1]-1, xyz[2]+1})
		}
		if xyz[1] < len(terrain[0])-1 && terrain[xyz[0]][xyz[1]+1] <= thisTerrain+1 {
			queue = append(queue, [3]int{xyz[0], xyz[1]+1, xyz[2]+1})
		}
	}

	return -1
}

func runPart2(terrain [][]byte, finish[2]int) (int) {
	queue := [][3]int{{finish[0], finish[1], 0}}
    visited := make([]bool, len(terrain)*len(terrain[0]))

	for len(queue) > 0 {
		var xyz [3]int
		xyz, queue = queue[0], queue[1:]

		// Loop prevention
		index := xyz[0]*len(terrain[0]) + xyz[1]
		if visited[index] {
			continue
		}
		visited[index] = true

		thisTerrain := terrain[xyz[0]][xyz[1]]

		if thisTerrain == 'a' {
			return xyz[2]
		}

		if xyz[0] > 0 && terrain[xyz[0]-1][xyz[1]]+1 >= thisTerrain {
			queue = append(queue, [3]int{xyz[0]-1, xyz[1], xyz[2]+1})
		}
		if xyz[0] < len(terrain)-1 && terrain[xyz[0]+1][xyz[1]]+1 >= thisTerrain {
			queue = append(queue, [3]int{xyz[0]+1, xyz[1], xyz[2]+1})
		}
		if xyz[1] > 0 && terrain[xyz[0]][xyz[1]-1]+1 >= thisTerrain {
			queue = append(queue, [3]int{xyz[0], xyz[1]-1, xyz[2]+1})
		}
		if xyz[1] < len(terrain[0])-1 && terrain[xyz[0]][xyz[1]+1]+1 >= thisTerrain {
			queue = append(queue, [3]int{xyz[0], xyz[1]+1, xyz[2]+1})
		}
	}

	return -1
}

func runDay12(inputFile string) (int, int) {
	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Printf("%v\n", err);
		os.Exit(1)
	}
	defer file.Close()

	scan := bufio.NewScanner(file)

	var terrain [][]byte
	var start [2]int
	var finish [2]int

	row := 0
	for scan.Scan() {
		var line []byte
		for col, c := range scan.Text() {
			if c == 'S' {
				line = append(line, 'a')
				start = [2]int{row, col}
			} else if c == 'E' {
				line = append(line, 'z')
				finish = [2]int{row, col}
			} else {
				line = append(line, byte(c))
			}
		}
		terrain = append(terrain, line)
		row++
	}

	part1 := runPart1(terrain, start, finish)
	part2 := runPart2(terrain, finish)

    return part1, part2
}

func main() {
	inputFile := flag.String("input", "day12-input.txt", "Problem input file")
	flag.Parse()

    part1, part2 := runDay12(*inputFile)
	fmt.Printf("%d\n%d\n", part1, part2)
}
