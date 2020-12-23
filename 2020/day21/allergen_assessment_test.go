package main

import "testing"

func Test(t *testing.T) {
	foods := parseInput("test_input")
	n, s := count(foods)

	if n != 5 {
		t.Errorf("Part one should be 5, but be %d", n)
	}

	if s != "mxmxvkd,sqjhc,fvjkl" {
		t.Errorf("Part two should be %s, but be %s", "mxmxvkd,sqjhc,fvjkl", s)
	}
}
