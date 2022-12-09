package main

import "testing"

func TestDay01(t *testing.T) {
	part1, part2 := runDay01("day01-test.txt")

	if part1 != 24000 {
		t.Errorf("Part 1 test returned %d; want 24000", part1)
	}

	if part2 != 45000 {
		t.Errorf("Part 2 test returned %d; want 45000", part2)
	}
}
