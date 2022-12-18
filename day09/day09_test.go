package main

import "testing"

func TestDay09(t *testing.T) {
	part1, part2 := runDay09("day09-test.txt")

	if part1 != 13 {
		t.Errorf("Part 1 test returned %d; want 13", part1)
	}

	if part2 != 1 {
		t.Errorf("Part 2 test 1 returned %d; want 1", part2)
	}

	_, part2 = runDay09("day09-test2.txt")
	if part2 != 36 {
		t.Errorf("Part 2 test 2 returned %d; want 36", part2)
	}
}
