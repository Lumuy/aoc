package main

import "testing"

func Test(t *testing.T) {
	input := parseInput("test_input")
	p1 := countFlashes(input, 100)
	p2 := countFlashes(input, -1)

	if p1 != 1656 {
		t.Errorf("Part one should be 1656. but be %d", p1)
	}

	if p2 != 195 {
		t.Errorf("Part two should be 195. but be %d", p2)
	}
}
