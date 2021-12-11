package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

var xlen, ylen int

type Point struct {
	x, y int
}

func parseInput(filename string) map[Point]int {
	m := make(map[Point]int)
	re := regexp.MustCompile(`\d`)
	dat, _ := ioutil.ReadFile(filename)
	for y, line := range strings.Split(strings.TrimSpace(string(dat)), "\n") {
		ylen++
		if xlen == 0 {
			xlen = len(line)
		}

		for x, str := range re.FindAllString(line, -1) {
			val, _ := strconv.Atoi(str)
			m[Point{x, y}] = val + 1
		}
	}

	return m
}

func (p Point) Add(q Point) Point {
	return Point{p.x + q.x, p.y + q.y}
}

func (p Point) Adjacents() (r []Point) {
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			if dx == 0 && dy == 0 {
				continue
			}
			x, y := p.x+dx, p.y+dy
			if x > -1 && x < xlen && y > -1 && y < ylen {
				r = append(r, Point{x, y})
			}
		}
	}

	return r
}

func cond(s, steps int) bool {
	if steps == -1 {
		return true
	}
	return s < steps
}

func countFlashes(m map[Point]int, steps int) (r int) {
	for s := 0; cond(s, steps); s++ {
		sn := 0
		nm := make(map[Point]int)
		var flashes []Point

		for p, v := range m {
			if v+1 == 11 {
				nm[p] = 1
				sn++
				flashes = append(flashes, p)
			} else {
				nm[p] = v + 1
			}
		}

		for len(flashes) != 0 {
			var tmp []Point
			for _, p := range flashes {
				for _, ap := range p.Adjacents() {
					if nm[ap] != 1 {
						if nm[ap]+1 == 11 {
							nm[ap] = 1
							sn++
							tmp = append(tmp, ap)
						} else {
							nm[ap] = nm[ap] + 1
						}
					}
				}
			}
			flashes = tmp
		}

		if sn == len(m) {
			return s + 1
		}
		r += sn
		m = nm
	}

	return r
}

func main() {
	input := parseInput("input")
	p1 := countFlashes(input, 100)
	p2 := countFlashes(input, -1)

	fmt.Println(p1)
	fmt.Println(p2)
}
