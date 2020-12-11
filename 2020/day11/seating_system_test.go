package main

import "testing"

func Test(t *testing.T) {
	seats := parseSeats("test_input")
	numbr := countOccupiedSeats(seats, 1)
	count := countOccupiedSeats(seats, 2)

	if numbr != 37 {
		t.Errorf("The count of occupied seats of part one should be 37, but be %d", numbr)
	}

	if count != 26 {
		t.Errorf("The count of occupied seats of part two should be 26, but be %d", count)
	}
}
