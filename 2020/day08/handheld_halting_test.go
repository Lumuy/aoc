package main

import "testing"

func Test(t *testing.T) {
	ins := getInstructions("test_input")
	cnt, _ := getAccumulatorValue(ins)
	val := getProgramValue(ins)

	if cnt != 5 {
		t.Errorf("The value of accumulator should be 5, but be %d", cnt)
	}

	if val != 8 {
		t.Errorf("The value of terminated program should be 8, but be %d", val)
	}
}
