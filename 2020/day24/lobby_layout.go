package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

var directions = map[string]Point{
	"e":  {2, 0},
	"se": {1, -3},
	"sw": {-1, -3},
	"w":  {-2, 0},
	"nw": {-1, 3},
	"ne": {1, 3},
}

type Point struct {
	x, y int
}

func (p Point) Add(q Point) Point {
	return Point{p.x + q.x, p.y + q.y}
}

func (p Point) Neibors() (r []Point) {
	for _, q := range directions {
		r = append(r, p.Add(q))
	}
	return r
}

func parseInput(filename string) (r [][]string) {
	dat, _ := ioutil.ReadFile(filename)
	re := regexp.MustCompile(`(?i:(se)|(e)|(sw)|(w)|(nw)|(ne))`)
	for _, line := range strings.Split(strings.TrimSpace(string(dat)), "\n") {
		t := []string{}
		for _, arr := range re.FindAllStringSubmatch(line, -1) {
			t = append(t, arr[0])
		}
		r = append(r, t)
	}
	return r
}

func count(m map[Point]bool, p Point) (r int) {
	for _, np := range p.Neibors() {
		if m[np] {
			r++
		}
	}
	return r
}

func run(ts [][]string) (nr, nc int) {
	m := map[Point]bool{}

	for _, t := range ts {
		p := Point{}
		for _, d := range t {
			p = p.Add(directions[d])
		}
		m[p] = !m[p]
	}

	for _, t := range m {
		if t {
			nr++
		}
	}

	for day := 1; day <= 100; day++ {
		nm := map[Point]bool{}
		ns := map[Point]int{}

		for p, t := range m {
			if t {
				for _, np := range p.Neibors() {
					ns[np] = ns[np] + 1
				}
			}
		}

		for p, c := range ns {
			if c == 1 && m[p] || c == 2 {
				nm[p] = true
			}
		}

		m = nm
	}

	for _, t := range m {
		if t {
			nc++
		}
	}

	return nr, nc
}

func main() {
	tiles := parseInput("input")
	fmt.Println(run(tiles))
}
