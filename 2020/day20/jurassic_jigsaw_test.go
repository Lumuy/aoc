package main

import "testing"

func Test(t *testing.T) {
	tiles := parseInput("test_input")
	p1, p2 := run(tiles)

	if p1 != 20899048083289 {
		t.Errorf("Part one should be 20899048083289, but be %d", p1)
	}

	if p2 != 273 {
		t.Errorf("Part two should be 273, but be %d", p2)
	}
}
