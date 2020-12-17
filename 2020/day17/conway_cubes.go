package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Point [4]int

func (p Point) Add(q Point) Point {
	for i := range p {
		p[i] += q[i]
	}
	return p
}

func parseCubes(filename string) map[Point]struct{} {
	m := make(map[Point]struct{})
	dat, _ := ioutil.ReadFile(filename)
	for y, s := range strings.Fields(string(dat)) {
		for x, r := range s {
			if r == '#' {
				m[Point{x, y}] = struct{}{}
			}
		}
	}
	return m
}

func delta(dim int) (ds []Point) {
	if dim == 0 {
		return []Point{{}}
	}
	for _, dm := range []int{0, 1, -1} {
		for _, p := range delta(dim - 1) {
			p[dim-1] = dm
			ds = append(ds, p)
		}
	}
	return
}

func countActiveCubes(area map[Point]struct{}, dim, cycles int) int {
	delta := delta(dim)[1:]

	for i := 0; i < cycles; i++ {
		neighbors := map[Point]int{}
		for p := range area {
			for _, d := range delta {
				neighbors[p.Add(d)]++
			}
		}

		na := make(map[Point]struct{})
		for p, n := range neighbors {
			if _, ok := area[p]; ok && n == 2 || n == 3 {
				na[p] = struct{}{}
			}
		}
		area = na
	}

	return len(area)
}

func main() {
	cubes := parseCubes("input")
	p1 := countActiveCubes(cubes, 3, 6)
	p2 := countActiveCubes(cubes, 4, 6)
	fmt.Println(p1, p2)
}
