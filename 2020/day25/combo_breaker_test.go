package main

import "testing"

func Test(t *testing.T) {
	pks := parseInput("test_input")
	num := run(pks)

	if num != 14897079 {
		t.Errorf("Part one should be 14897079, but be %d", num)
	}
}
