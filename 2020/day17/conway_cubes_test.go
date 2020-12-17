package main

import "testing"

func Test(t *testing.T) {
	cubes := parseCubes("test_input")
	cnt1 := countActiveCubes(cubes, 3, 6)
	cnt2 := countActiveCubes(cubes, 4, 6)

	if cnt1 != 112 {
		t.Errorf("Active cubes of part one should be 112, but be %d", cnt1)
	}

	if cnt2 != 848 {
		t.Errorf("Active cubes of part two should be 848, but be %d", cnt2)
	}
}
