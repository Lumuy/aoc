package main

import "testing"

func Test(t *testing.T) {
	parseInput("test_input")
	count1 := countMessagesMatched("0", 1)
	count2 := countMessagesMatched("0", 2)

	if count1 != 3 {
		t.Errorf("Part one should be 3, but be %d", count1)
	}

	if count2 != 12 {
		t.Errorf("Part two should be 3, but be %d", count2)
	}
}
