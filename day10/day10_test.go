package main

import "testing"

func TestDay10(t *testing.T) {
	part1, part2 := runDay10("day10-test.txt")

	if part1 != 13140 {
		t.Errorf("Part 1 test returned %d; want 13140", part1)
	}

	part2output := `##..##..##..##..##..##..##..##..##..##..
###...###...###...###...###...###...###.
####....####....####....####....####....
#####.....#####.....#####.....#####.....
######......######......######......####
#######.......#######.......#######.....
`
	if part2 != part2output {
		t.Errorf("Part 2 test 1 returned unexpected output: \n%s\nExpected:\n%s", part2, part2output)
	}
}
