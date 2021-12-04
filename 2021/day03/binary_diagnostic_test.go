package main

import "testing"

func Test(t *testing.T) {
	input := parseInput("test_input")
	p1 := countConsumption(input)
	p2 := countSupportRating(input)

	if p1 != 198 {
		t.Errorf("Part one should be 198, but be %d", p1)
	}

	if p2 != 230 {
		t.Errorf("Part two should be 230, but be %d", p2)
	}
}
