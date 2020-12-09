package main

import "testing"

func Test(t *testing.T) {
	numbers := parseNumbers("test_input")
	num := getFirstNumberUnfollowRule(numbers, 5)
	wks := getEncryptionWeakness(numbers, num)

	if num != 127 {
		t.Errorf("The first number not follow this rule should be 127, but be %d", num)
	}

	if wks != 62 {
		t.Errorf("The encryption weakness should be 62, but be %d", wks)
	}
}
