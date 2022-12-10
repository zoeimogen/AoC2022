package main

import "testing"

func TestDay05(t *testing.T) {
	part1, part2 := runDay05("day05-test.txt")

	if part1 != "CMZ" {
		t.Errorf("Part 1 test returned %s; want CMZ", part1)
	}

	if part2 != "MCD" {
		t.Errorf("Part 2 test returned %s; want MCD", part2)
	}
}
