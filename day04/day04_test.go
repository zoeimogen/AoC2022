package main

import "testing"

func TestDay04(t *testing.T) {
	part1, part2 := runDay04("day04-test.txt")

	if part1 != 2 {
		t.Errorf("Part 1 test returned %d; want 2", part1)
	}

	if part2 != 4 {
		t.Errorf("Part 2 test returned %d; want 4", part2)
	}
}
