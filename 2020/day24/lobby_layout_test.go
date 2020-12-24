package main

import "testing"

func Test(t *testing.T) {
	tiles := parseInput("test_input")
	p1, p2 := run(tiles)

	if p1 != 10 {
		t.Errorf("Part one should be 10, but be  %d", p1)
	}

	if p2 != 2208 {
		t.Errorf("Part two should be 2208, but be  %d", p2)
	}
}
