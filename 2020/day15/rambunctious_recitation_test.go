package main

import "testing"

func Test(t *testing.T) {
	input := parseInput("test_input")
	numb1 := getSpokenNumber(input, 2020)
	numb2 := getSpokenNumber(input, 30000000)

	if numb1 != 436 {
		t.Errorf("The 2020th number spoken should be 436, but be %d", numb1)
	}

	if numb2 != 175594 {
		t.Errorf("The 30000000th number spoken should be 175594, but be %d", numb2)
	}
}
