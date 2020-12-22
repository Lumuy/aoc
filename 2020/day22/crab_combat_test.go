package main

import "testing"

func Test(t *testing.T) {
	decks := parseInput("test_input")
	_, score1 := play(decks, 1)
	_, score2 := play(decks, 2)

	if score1 != 306 {
		t.Errorf("The winning player's score of part one should be 306, but be %d", score1)
	}

	if score2 != 291 {
		t.Errorf("The winning player's score of part two should be 291, but be %d", score2)
	}
}
