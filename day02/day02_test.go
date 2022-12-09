package main

import "testing"

func TestDay02(t *testing.T) {
	part1, part2 := runDay02("day02-test.txt")

	if part1 != 15 {
		t.Errorf("Part 1 test returned %d; want 15", part1)
	}

	if part2 != 12 {
		t.Errorf("Part 2 test returned %d; want 12", part2)
	}
}
