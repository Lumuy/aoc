package main

import "testing"

func Test(t *testing.T) {
	input := parseInput("test_input")
	p1 := count(input, false)
	p2 := count(input, true)

	if p1 != 5 {
		t.Errorf("Part one should be 5, but be %d", p1)
	}

	if p2 != 12 {
		t.Errorf("Part two should be 12, but be %d", p2)
	}
}
