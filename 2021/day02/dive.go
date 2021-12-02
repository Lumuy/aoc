package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Point struct {
	x, y int
}

func (p Point) Add(q Point) Point {
	return Point{p.x + q.x, p.y + q.y}
}

func parseInput(filename string) (r []Point) {
	dat, _ := ioutil.ReadFile(filename)
	for _, line := range strings.Split(strings.TrimSpace(string(dat)), "\n") {
		arr := strings.Split(line, " ")
		dep, _ := strconv.Atoi(arr[1])

		switch arr[0] {
		case "forward":
			r = append(r, Point{dep, 0})
		case "down":
			r = append(r, Point{0, dep})
		case "up":
			r = append(r, Point{0, -dep})
		}
	}

	return r
}

func calc_multiply(instructions []Point, part int) (r int) {
	var submarin, aim Point
	for _, p := range instructions {
		switch part {
		case 1:
			submarin = submarin.Add(p)
		case 2:
			if p.x != 0 {
				submarin = Point{submarin.x + p.x, submarin.y + aim.y * p.x}
			} else {
				aim = aim.Add(p)
			}
		}
	}

	return submarin.x * submarin.y
}

func main() {
	input := parseInput("input")
	p1 := calc_multiply(input, 1)
	p2 := calc_multiply(input, 2)

	fmt.Println(p1)
	fmt.Println(p2)
}
