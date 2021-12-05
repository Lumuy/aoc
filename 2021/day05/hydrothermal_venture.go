package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type Point struct {
	x, y int
}

type Line struct {
	start, end Point
}

func parseInput(filename string) (r []Line) {
	dat, _ := ioutil.ReadFile(filename)
	re := regexp.MustCompile(`\d+`)
	for _, str := range strings.Split(strings.TrimSpace(string(dat)), "\n") {
		var arr [4]int
		for i, s := range re.FindAllString(str, -1) {
			n, _ := strconv.Atoi(s)
			arr[i] = n
		}
		r = append(r, Line{start: Point{arr[0], arr[1]}, end: Point{arr[2], arr[3]}})
	}

	return r
}

func compare(x, y int) (int, int) {
	if x < y {
		return x, y
	}
	return y, x
}

func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}

func count(lines []Line, diagonal bool) (r int) {
	m := make(map[Point]int)

	for _, line := range lines {
		if line.start.x == line.end.x {
			min, max := compare(line.start.y, line.end.y)
			for i := min; i <= max; i++ {
				p := Point{line.start.x, i}
				m[p] = m[p] + 1
			}
		}
		if line.start.y == line.end.y {
			min, max := compare(line.start.x, line.end.x)
			for i := min; i <= max; i++ {
				p := Point{i, line.start.y}
				m[p] = m[p] + 1
			}
		}

		lx, ly := line.end.x-line.start.x, line.end.y-line.start.y
		if diagonal && abs(lx) == abs(ly) {
			var dx, dy int
			if lx > 0 && ly > 0 {
				dx, dy = 1, 1
			}
			if lx > 0 && ly < 0 {
				dx, dy = 1, -1
			}
			if lx < 0 && ly > 0 {
				dx, dy = -1, 1
			}
			if lx < 0 && ly < 0 {
				dx, dy = -1, -1
			}

			for i := 0; i <= abs(lx); i++ {
				p := Point{line.start.x + i*dx, line.start.y + i*dy}
				m[p] = m[p] + 1
			}
		}
	}

	for _, v := range m {
		if v > 1 {
			r++
		}
	}

	return r
}

func main() {
	input := parseInput("input")
	p1 := count(input, false)
	p2 := count(input, true)

	fmt.Println(p1)
	fmt.Println(p2)
}
