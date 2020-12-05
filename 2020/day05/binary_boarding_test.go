package main

import "testing"

func Test(t *testing.T) {
	id1 := parsePassId("BFFFBBFRRR")
	id2 := parsePassId("FFFBBBFRRR")
	id3 := parsePassId("BBFFBBFRLL")

	if id1 != 567 {
		t.Errorf("Pass id should be 567, but be %d", id1)
	}
	if id2 != 119 {
		t.Errorf("Pass id should be 119, but be %d", id2)
	}
	if id3 != 820 {
		t.Errorf("Pass id should be 820, but be %d", id3)
	}
}
