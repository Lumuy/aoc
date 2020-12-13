package main

import "testing"

func Test(t *testing.T) {
	start, buses := parseBuses("test_input")
	num := getMultipliedNumber(start, buses)
	tim := getEarliestTimestamp(buses)

	if num != 295 {
		t.Errorf("The multiplied number should be 295, but be %d", num)
	}

	if tim != 1068781 {
		t.Errorf("The earliest timestamp should be 1068781, but be %d", tim)
	}
}
