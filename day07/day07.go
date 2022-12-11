package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

type Entry struct {
	size int
	dir  *Directory
}

type Directory struct {
	parent  *Directory
	entries map[string]Entry
}

func dirSize(d Directory) int {
	var total int

	for _, e := range d.entries {
		if e.dir != nil {
			total += dirSize(*e.dir)
		} else {
			total += e.size
		}
	}

	return total
}

func dirSizePart1(d Directory) int {
	var total int

	for _, e := range d.entries {
		if e.dir != nil {
			total += dirSizePart1(*e.dir)
		}
	}

	t := dirSize(d)
	if t <= 100000 {
		total += t
	}

	return total
}

func part2(root Directory) int {
	const fsSize = 70000000
	const updateSize = 30000000

	spaceUsed := dirSize(root)
	spaceAvailable := fsSize - spaceUsed
	spaceNeeded := updateSize - spaceAvailable

	return dirSizePart2(root, spaceNeeded)
}

func dirSizePart2(d Directory, needed int) int {
	var s int

	best := 0
	for _, e := range d.entries {
		if e.dir != nil {
			s = dirSizePart2(*e.dir, needed)
			if (s < best || best == 0) && s >= needed {
				best = s
			}
		}
	}

	s = dirSize(d)
	if (s < best || best == 0) && s >= needed {
		best = s
	}

	return best
}

func loadData(file io.Reader) *Directory {
	root := new(Directory)
	root.entries = make(map[string]Entry)
	cwd := root

	scan := bufio.NewScanner(file)
	for scan.Scan() {
		if scan.Text()[0:4] == "$ cd" {
			// Command "cd"
			newdir := scan.Text()[5:]
			if newdir == "/" {
				cwd = root
			} else if newdir == ".." {
				cwd = cwd.parent
			} else {
				d, ok := (*cwd).entries[newdir]
				if ok {
					cwd = d.dir
				} else {
					// Create a new directory
					newdirnode := new(Directory)
					newdirnode.parent = cwd
					newdirnode.entries = make(map[string]Entry)
					cwd = newdirnode
				}
			}
		} else if scan.Text()[0] != byte('$') {
			// All non-commands
			if scan.Text()[0:4] == "dir " {
				// Create a new directory
				newdir := scan.Text()[4:]
				newdirnode := new(Directory)
				newdirnode.parent = cwd
				newdirnode.entries = make(map[string]Entry)
				cwd.entries[newdir] = Entry{0, newdirnode}
			} else {
				var size int
				var name string
				v, _ := fmt.Sscanf(scan.Text(), "%d %s", &size, &name)
				if v != 2 {
					fmt.Printf("Can't parse entry: %s\n", scan.Text())
				} else {
					cwd.entries[name] = Entry{size, nil}
				}
			}
		}
	}

	return root
}

func runDay07(inputFile string) (int, int) {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer file.Close()

	root := loadData(file)

	return dirSizePart1(*root), part2(*root)
}

func main() {
	var inputFile = flag.String("input", "day07-input.txt", "Problem input file")
	flag.Parse()
	part1, part2 := runDay07(*inputFile)
	fmt.Printf("%d %d\n", part1, part2)
}
