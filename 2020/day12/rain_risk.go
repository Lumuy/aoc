package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type instruction struct {
	action rune
	value  int
}

func parseInstructions(filename string) (r []instruction) {
	dat, _ := ioutil.ReadFile(filename)
	for _, line := range strings.Split(strings.TrimSpace(string(dat)), "\n") {
		n, _ := strconv.Atoi(line[1:])
		r = append(r, instruction{action: []rune(line)[0], value: n})
	}
	return r
}

func move(action rune, val, x, y int) (int, int) {
	switch action {
	case 'E':
		x += val
	case 'N':
		y += val
	case 'W':
		x -= val
	case 'S':
		y -= val
	}
	return x, y
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// counter-clockwise
func rotateWaypoint(step, wx, wy int) (rx, ry int) {
	switch step {
	case 1:
		rx = -wy
		ry = wx
	case 2:
		rx = -wx
		ry = -wy
	case 3:
		rx = wy
		ry = -wx
	}
	return rx, ry
}

func getManhattanDistance(instructions []instruction, part int) int {
	var x, y, facing_idx int
	directions := []rune{'E', 'N', 'W', 'S'}
	wx, wy := 10, 1

	for _, ins := range instructions {
		switch ins.action {
		case 'E', 'N', 'W', 'S':
			if part == 1 {
				x, y = move(ins.action, ins.value, x, y)
			}
			if part == 2 {
				wx, wy = move(ins.action, ins.value, wx, wy)
			}
		case 'L', 'R':
			var step int
			if ins.action == 'L' {
				step = ins.value / 90
			}
			if ins.action == 'R' {
				step = (360 - ins.value) / 90
			}

			if part == 1 {
				facing_idx = (facing_idx + step) % 4
			}
			if part == 2 {
				wx, wy = rotateWaypoint(step, wx, wy)
			}
		case 'F':
			if part == 1 {
				x, y = move(directions[facing_idx], ins.value, x, y)
			}
			if part == 2 {
				x += ins.value * wx
				y += ins.value * wy
			}
		}
	}

	return abs(x) + abs(y)
}

func main() {
	ins := parseInstructions("input")
	dis := getManhattanDistance(ins, 1)
	dst := getManhattanDistance(ins, 2)
	fmt.Println(dis, dst)
}
