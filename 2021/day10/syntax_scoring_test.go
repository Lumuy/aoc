package main

import "testing"

func Test(t *testing.T) {
	input := parseInput("test_input")
	p1, p2 := countScores(input)

	if p1 != 26397 {
		t.Errorf("Part one should be 26397. but be %d", p1)
	}

	if p2 != 288957 {
		t.Errorf("Part two should be 288957. but be %d", p2)
	}
}
