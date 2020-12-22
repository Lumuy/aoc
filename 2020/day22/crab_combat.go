package main

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"
)

func parseInput(filename string) (decks [][]int) {
	dat, _ := ioutil.ReadFile(filename)
	for _, content := range strings.Split(strings.TrimSpace(string(dat)), "\n\n") {
		var deck []int
		for _, line := range strings.Split(strings.TrimSpace(content), "\n")[1:] {
			score, _ := strconv.Atoi(line)
			deck = append(deck, score)
		}
		decks = append(decks, deck)
	}
	return decks
}

func copyDecks(decks [][]int) (r [][]int) {
	for _, arr := range decks {
		var add []int
		for _, v := range arr {
			add = append(add, v)
		}
		r = append(r, add)
	}
	return r
}

func contains(records [][]int, arr []int) bool {
	for _, v := range records {
		if reflect.DeepEqual(v, arr) {
			return true
		}
	}
	return false
}

func left(arr []int) []int {
	if len(arr) > 1 {
		return arr[1:]
	}
	return []int{}
}

func play(decks [][]int, part int) (win, sum int) {
	records0, records1 := [][]int{}, [][]int{}
	decks = copyDecks(decks)

	for len(decks[0])*len(decks[1]) != 0 {
		if part == 2 {
			if contains(records0, decks[0]) || contains(records1, decks[1]) {
				return 0, 0
			}
			records0 = append(records0, decks[0])
			records1 = append(records1, decks[1])
		}

		if part == 2 && decks[0][0] <= len(decks[0][1:]) && decks[1][0] <= len(decks[1][1:]) {
			twin, _ := play([][]int{decks[0][1 : decks[0][0]+1], decks[1][1 : decks[1][0]+1]}, part)
			if twin == 0 {
				decks[0] = append(decks[0][1:], []int{decks[0][0], decks[1][0]}...)
				decks[1] = decks[1][1:]
			}
			if twin == 1 {
				decks[1] = append(decks[1][1:], []int{decks[1][0], decks[0][0]}...)
				decks[0] = decks[0][1:]
			}
		} else {
			if decks[0][0] > decks[1][0] {
				decks[0] = append(left(decks[0]), []int{decks[0][0], decks[1][0]}...)
				decks[1] = left(decks[1])
			} else {
				decks[1] = append(left(decks[1]), []int{decks[1][0], decks[0][0]}...)
				decks[0] = left(decks[0])
			}
		}
	}

	if len(decks[0]) == 0 {
		win = 1
	} else {
		win = 0
	}

	for i, n := range decks[win] {
		sum += n * (len(decks[win]) - i)
	}

	return win, sum
}

func main() {
	decks := parseInput("input")
	_, score1 := play(decks, 1)
	_, score2 := play(decks, 2)
	fmt.Println(score1, score2)
}
