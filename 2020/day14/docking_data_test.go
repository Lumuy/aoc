package main

import "testing"

func Test(t *testing.T) {
	program1 := parseProgram("test_input")
	program2 := parseProgram("test_input_two")
	sum1 := getMemorySum(program1, 1)
	sum2 := getMemorySum(program2, 2)

	if sum1 != 165 {
		t.Errorf("The sum of all values left in memory of part one should be 165, but be %d", sum1)
	}

	if sum2 != 208 {
		t.Errorf("The sum of all values left in memory of part two should be 208, but be %d", sum2)
	}
}
