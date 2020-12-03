package main

import "testing"

func Test(t *testing.T) {
	area := parseInput("test_input")
	count1 := countTrees(area, 3, 1)
	count2 := getMultiply(area)

	if count1 != 7 {
		t.Errorf("Part one valid tree number should be 7, but be %d", count1)
	}

	if count2 != 336 {
		t.Errorf("Part two valid tree number should be 336, but be %d", count2)
	}
}
