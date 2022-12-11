package main

import "testing"

func TestDay08(t *testing.T) {
	part1, part2 := runDay08("day08-test.txt")

	if part1 != 21 {
		t.Errorf("Part 1 test returned %d; want 21", part1)
	}

	if part2 != 8 {
		t.Errorf("Part 2 test returned %d; want 8", part2)
	}
}
