package main

import "testing"

func Test(t *testing.T) {
	input := parseInput("test_input")
	p1 := count(input, 80)
	p2 := count(input, 256)

	if p1 != 5934 {
		t.Errorf("Part one should be 5934, but be %d", p1)
	}

	if p2 != 26984457539 {
		t.Errorf("Part one should be 26984457539, but be %d", p2)
	}
}
