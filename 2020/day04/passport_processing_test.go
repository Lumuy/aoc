package main

import "testing"

func Test(t *testing.T) {
	passports := getPassports("test_input")
	count := countValidPassports(passports)

	if count != 2 {
		t.Errorf("Part one valid passports number should be 2, but be %d", count)
	}

	for _, pt := range getPassports("invalid_input") {
		if pt.isStrictValid() {
			t.Error("Passport hould be invalid")
		}
	}

	for _, pt := range getPassports("valid_input") {
		if !pt.isStrictValid() {
			t.Error("Passport hould be valid")
		}
	}
}
