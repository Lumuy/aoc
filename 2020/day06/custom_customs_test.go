package main

import "testing"

func Test(t *testing.T) {
	c1, c2 := getSumOfCounts(parseAnswers("test_input"))

	if c1 != 11 {
		t.Errorf("Sum of counts for anyone should be 11, but be %d", c1)
	}

	if c2 != 6 {
		t.Errorf("Sum of counts for everyone should be 6, but be %d", c2)
	}
}
