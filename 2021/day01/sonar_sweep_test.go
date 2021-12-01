package main

import "testing"

func Test(t *testing.T) {
	arr1 := parseInput("test_input")
	arr2 := three_measurement(arr1)
	p1 := count(arr1)
	p2 := count(arr2)

	if p1 != 7 {
		t.Errorf("Part one should be 7, but be %d", p1)
	}

	if p2 != 5 {
		t.Errorf("Part one should be 5, but be %d", p2)
	}
}
