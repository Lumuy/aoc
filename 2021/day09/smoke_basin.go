package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

var xlen, ylen int

type Point struct {
	x, y int
}

func parseInput(filename string) map[Point]int {
	r := make(map[Point]int)
	dat, _ := ioutil.ReadFile(filename)
	re := regexp.MustCompile(`\d`)
	for y, line := range strings.Split(strings.TrimSpace(string(dat)), "\n") {
		ylen++
		if xlen == 0 {
			xlen = len(line)
		}

		for x, s := range re.FindAllString(line, -1) {
			n, _ := strconv.Atoi(s)
			r[Point{x, y}] = n + 1
		}
	}

	return r
}

func adjacentPoints(p Point, counted map[Point]bool, m map[Point]int) (r []Point) {
	for _, ap := range []Point{{x: p.x, y: p.y - 1}, {x: p.x, y: p.y + 1}, {x: p.x - 1, y: p.y}, {x: p.x + 1, y: p.y}} {
		if counted[ap] == false && m[ap] > 0 && m[ap] < 10 {
			r = append(r, ap)
		}
	}
	return r
}

func countRisks(m map[Point]int) (r, s int) {
	var sizes []int

	for y := 0; y < ylen; y++ {
		for x := 0; x < xlen; x++ {
			up := m[Point{x, y}] < m[Point{x, y - 1}] || m[Point{x, y - 1}] == 0
			dn := m[Point{x, y}] < m[Point{x, y + 1}] || m[Point{x, y + 1}] == 0
			lt := m[Point{x, y}] < m[Point{x - 1, y}] || m[Point{x - 1, y}] == 0
			rt := m[Point{x, y}] < m[Point{x + 1, y}] || m[Point{x + 1, y}] == 0

			if up && dn && lt && rt {
				r += m[Point{x, y}]
				sizes = append(sizes, countBasin(Point{x, y}, m))
			}
		}
	}

	sort.Slice(sizes, func(i, j int) bool {
		return sizes[i] > sizes[j]
	})
	s = sizes[0] * sizes[1] * sizes[2]

	return r, s
}

func countBasin(p Point, m map[Point]int) int {
	counted := make(map[Point]bool)
	counted[p] = true
	var points []Point
	points = append(points, p)

	for len(points) != 0 {
		var tmp []Point
		for _, sp := range points {
			for _, ap := range adjacentPoints(sp, counted, m) {
				tmp = append(tmp, ap)
				counted[ap] = true
			}
		}
		points = tmp
	}

	return len(counted)
}

func main() {
	input := parseInput("input")
	p1, p2 := countRisks(input)

	fmt.Println(p1)
	fmt.Println(p2)
}
