package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func recursiveCheck(input []byte) bool {
	if len(input) == 1 {
		return true
	}

	for i := 1; i < len(input); i++ {
		if input[0] == input[i] {
			return false
		}
	}

	return recursiveCheck(input[1:])
}

func findMarker(length int, input []byte) int {
	for i := length; i < len(input); i++ {
		if recursiveCheck(input[i-length : i]) {
			return i
		}
	}

	println("Can't find marker")
	return 0
}

func main() {
	var inputFile = flag.String("input", "day06-input.txt", "Problem input file")
	flag.Parse()
	input, err := os.ReadFile(*inputFile)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	fmt.Printf("%d, %d\n", findMarker(4, input), findMarker(14, input))
}
