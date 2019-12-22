package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func dealIntoNewStack(cards []int) (r []int) {
	for i := len(cards) - 1; i >= 0; i-- {
		r = append(r, cards[i])
	}
	return r
}

func cutCards(cards []int, n int) (r []int) {
	if n > 0 {
		r = append(cards[n:], cards[0:n]...)
	} else if n < 0 {
		s := len(cards) + n
		r = append(cards[s:], cards[:s]...)
	}
	return r
}

func dealWithIncrement(cards []int, n int) (r []int) {
	size := len(cards)
	// Grow result slice size with enough space
	for i := 0; i < size; i++ {
		r = append(r, 0)
	}

	ri := 0
	for i := 0; i < size; i++ {
		r[ri] = cards[i]
		ri = (ri + n) % size
	}

	return r
}

func process(inputs string, size int) []int {
	cards := genCards(size)

	for _, line := range strings.Split(inputs, "\n") {
		if strings.Contains(line, "new") {
			cards = dealIntoNewStack(cards)
		} else if strings.Contains(line, "cut") {
			n := parseNumber(line)
			cards = cutCards(cards, n)
		} else if strings.Contains(line, "increment") {
			n := parseNumber(line)
			cards = dealWithIncrement(cards, n)
		}
	}

	return cards
}

func main() {
	inputs := readFile("input")
	cards := process(inputs, 10007)

	// Part 1
	for i, v := range cards {
		if v == 2019 {
			fmt.Println(i)
			break
		}
	}

	// Part 2
}

func parseNumber(s string) int {
	re := regexp.MustCompile(`[-]?\d+`)
	rs := re.FindString(s)
	n, err := strconv.Atoi(rs)
	check(err)

	return n
}

func readFile(filename string) string {
	bytes, err := ioutil.ReadFile(filename)
	check(err)
	return strings.TrimSpace(string(bytes))
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func genCards(size int) (r []int) {
	for i := 0; i < size; i++ {
		r = append(r, i)
	}
	return r
}
