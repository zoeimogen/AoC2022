package main

import "testing"

func TestDay03(t *testing.T) {
	part1, part2 := runDay03("day03-test.txt")

	if part1 != 157 {
		t.Errorf("Part 1 test returned %d; want 157", part1)
	}

	if part2 != 70 {
		t.Errorf("Part 2 test returned %d; want 70", part2)
	}
}
