package main

import "testing"

func Test(t *testing.T) {
	input := parseInput("test_input")
	p1, p2 := countRisks(input)

	if p1 != 15 {
		t.Errorf("Part one should be 15. but be %d", p1)
	}

	if p2 != 1134 {
		t.Errorf("Part two should be 1134. but be %d", p2)
	}
}
