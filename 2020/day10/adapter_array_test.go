package main

import "testing"

func Test(t *testing.T) {
	jolts := parseJolts("test_input")
	numbr := getMultipliedNum(jolts)
	count := countArrangements(jolts)

	if numbr != 220 {
		t.Errorf("The multiplied number of 1-jolt and 3-jolt should be 220, but be %d", numbr)
	}

	if count != 19208 {
		t.Errorf("The total number of arrangement for adapters should be 19208, but be %d", count)
	}
}
