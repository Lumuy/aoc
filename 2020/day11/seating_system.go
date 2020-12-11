package main

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"strings"
)

const floor = '.'
const empty = 'L'
const occup = '#'

type seat struct {
	x, y int
}

func searchDirectionSeats(seats map[seat]rune, s seat, dx, dy int) (r int) {
	st := seat{x: s.x, y: s.y}
	for {
		st = seat{x: st.x + dx, y: st.y + dy}
		if seats[st] != floor {
			if seats[st] == occup {
				r++
			}
			break
		}
	}
	return r
}

func occupiedAdjacentSeats(s seat, seats map[seat]rune, part int) (r int) {
	for _, arr := range [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}, {1, -1}, {-1, -1}, {1, 1}, {-1, 1}} {
		switch part {
		case 1:
			if seats[seat{x: s.x + arr[0], y: s.y + arr[1]}] == occup {
				r++
			}
		case 2:
			r += searchDirectionSeats(seats, s, arr[0], arr[1])
		}
	}
	return r
}

func parseSeats(filename string) map[seat]rune {
	seats := make(map[seat]rune)
	dat, _ := ioutil.ReadFile(filename)
	for y, line := range strings.Split(strings.TrimSpace(string(dat)), "\n") {
		for x, s := range []rune(line) {
			seats[seat{x: x, y: y}] = s
		}
	}
	return seats
}

func compareNum(part int) (r int) {
	switch part {
	case 1:
		r = 4
	case 2:
		r = 5
	}
	return r
}

func copySeats(source map[seat]rune) map[seat]rune {
	r := make(map[seat]rune)
	for k, v := range source {
		r[k] = v
	}
	return r
}

func countOccupiedSeats(seats map[seat]rune, part int) (r int) {
	changed := make(map[seat]rune)
	for {
		for seat, status := range seats {
			switch status {
			case empty:
				if occupiedAdjacentSeats(seat, seats, part) == 0 {
					changed[seat] = occup
				}
			case occup:
				if occupiedAdjacentSeats(seat, seats, part) >= compareNum(part) {
					changed[seat] = empty
				}
			case floor:
				changed[seat] = floor
			}
		}
		if reflect.DeepEqual(seats, changed) {
			for _, s := range seats {
				if s == occup {
					r++
				}
			}
			break
		} else {
			seats = copySeats(changed)
		}
	}
	return r
}

func main() {
	seats := parseSeats("input")
	numbr := countOccupiedSeats(seats, 1)
	count := countOccupiedSeats(seats, 2)
	fmt.Println(numbr, count)
}
