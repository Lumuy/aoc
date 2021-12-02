package main

import "testing"

func Test(t *testing.T) {
	input := parseInput("test_input")
  p1 := calc_multiply(input, 1)
  p2 := calc_multiply(input, 2)

	if p1 != 150 {
		t.Errorf("Part one should be 150, but be %d", p1)
	}

	if p2 != 900 {
		t.Errorf("Part two should be 900, but be %d", p2)
	}
}
