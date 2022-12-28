package main

import (
	"testing"
)

func TestDay12(t *testing.T) {
	part1, part2 := runDay12("day12-test.txt")

	if part1 != 31 {
		t.Errorf("Part 1 test returned %d; want 31", part1)
	}

	if part2 != 29 {
		t.Errorf("Part 2 test returned %d; want 29", part2)
	}
}
