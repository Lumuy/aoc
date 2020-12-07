package main

import "testing"

func Test(t *testing.T) {
	rules1 := parseRules("test_input")
	rules2 := parseRules("test_input_two")
	count1 := countBagColors(rules1, "shiny gold")
	count2 := countIndividualBags(rules1, "shiny gold")
	count3 := countIndividualBags(rules2, "shiny gold")

	if count1 != 4 {
		t.Errorf("Should be 4 bag colors, but be %d", count1)
	}

	if count2 != 32 {
		t.Errorf("Should be 32 bag colors, but be %d", count2)
	}

	if count3 != 126 {
		t.Errorf("Should be 126 bag colors, but be %d", count3)
	}
}
