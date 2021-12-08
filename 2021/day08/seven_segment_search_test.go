package main

import "testing"

func Test(t *testing.T) {
	input := parseInput("test_input")
	p1 := countUniqDigits(input)
	input2 := parseInput("test_input_two")
	p2 := sumOutputs(input2)

	if p1 != 26 {
		t.Errorf("Part one should be 26. but be %d", p1)
	}

	if p2 != 5353 {
		t.Errorf("Part one should be 5353. but be %d", p2)
	}
}
