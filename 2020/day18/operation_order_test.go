package main

import "testing"

func Test(t *testing.T) {
	inp := parseInput("test_input")
	cn1 := sum(inp, 1)
	cn2 := sum(inp, 2)

	if cn1 != 13632 {
		t.Errorf("Part one should be 13632, but be %d", cn1)
	}

	if cn2 != 23340 {
		t.Errorf("Part two should be 23340, but be %d", cn2)
	}
}
