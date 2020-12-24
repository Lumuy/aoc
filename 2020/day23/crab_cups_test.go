package main

import "testing"

func Test(t *testing.T) {
	labels := "389125467"
	str1 := play(labels, len(labels), 100)
	str2 := play(labels, 1000000, 10000000)

	if str1 != "67384529" {
		t.Errorf("Part one should be %s, but be %s", "67384529", str1)
	}

	if str2 != "149245887792" {
		t.Errorf("Part two should be %s, but be %s", "149245887792", str2)
	}
}
