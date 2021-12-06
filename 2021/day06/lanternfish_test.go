package main

import "testing"

func Test(t *testing.T) {
	input := parseInput("test_input")
	p1 := count(input, 80)

	if p1 != 5934 {
		t.Errorf("Part one should be 5934, but be %d", p1)
	}
}
