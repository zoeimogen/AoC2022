package main

import (
	"testing"
)

func TestDay06(t *testing.T) {
	testdata := [][]interface{}{
		{"mjqjpqmgbljsphdztnvjfqwrcgsmlb", 7, 19},
		{"bvwbjplbgvbhsrlpgdmjqwftvncz", 5, 23},
		{"nppdvjthqldpwncqszvftbrmjlhg", 6, 23},
		{"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 10, 29},
		{"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 11, 26},
	}

	for i, _ := range testdata {
		result := findMarker(4, []byte(testdata[i][0].(string)))
		if result != testdata[i][1].(int) {
			t.Errorf("Testing %s with marker length 4 returned %d; want %d", testdata[i][0].(string), result, testdata[i][1].(int))
		}

		result = findMarker(14, []byte(testdata[i][0].(string)))
		if result != testdata[i][2].(int) {
			t.Errorf("Testing %s with marker length 14 returned %d; want %d", testdata[i][0].(string), result, testdata[i][2].(int))
		}
	}
}
