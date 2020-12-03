package main

import "testing"

func Test(t *testing.T) {
	passwords := parseInput("test_input")
	count1 := countValidPasswords(passwords)
	count2 := countValidPasswordsTwo(passwords)

	if count1 != 2 {
		t.Errorf("Part one valid password number should be 2, but be %d", count1)
	}

	if count2 != 1 {
		t.Errorf("Part two valid password number should be 1, but be %d", count2)
	}
}
