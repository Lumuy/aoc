package main

import "testing"

func Test(t *testing.T) {
	ins := parseInstructions("test_input")
	dis := getManhattanDistance(ins, 1)
	dst := getManhattanDistance(ins, 2)

	if dis != 25 {
		t.Errorf("The ship's manhattan distance of part one should be 25, but be %d", dis)
	}

	if dst != 286 {
		t.Errorf("The ship's manhattan distance of part two should be 286, but be %d", dst)
	}
}
