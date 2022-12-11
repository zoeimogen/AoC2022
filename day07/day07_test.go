package main

import "testing"

func TestDay07(t *testing.T) {
	part1, part2 := runDay07("day07-test.txt")

	if part1 != 95437 {
		t.Errorf("Part 1 test returned %d; want 95437", part1)
	}

	if part2 != 24933642 {
		t.Errorf("Part 2 test returned %d; want 24933642", part2)
	}
}
