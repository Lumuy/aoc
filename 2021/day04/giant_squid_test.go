package main

import "testing"

func Test(t *testing.T) {
	arr, matrixs := parseInput("test_input")
	p1, p2 := countFinalScore(arr, matrixs)

	if p1 != 4512 {
		t.Errorf("Part one should be 4512, but be %d", p1)
	}

	if p2 != 1924 {
		t.Errorf("Part two should be 1924, but be %d", p1)
	}
}
