package main

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

type Point struct {
	x, y float64
}

func getSigntPoints(s, d Point) (r []Point) {
	var p1 Point
	var p2 Point

	if s.x > d.x {
		p2, p1 = s, d
	} else if s.x < d.x {
		p2, p1 = d, s
	} else {
		if s.y > d.y {
			p2, p1 = s, d
		} else {
			p2, p1 = d, s
		}
		ydis := p2.y - p1.y

		for dy := 1.0; dy < ydis; dy++ {
			r = append(r, Point{x: p1.x, y: p1.y + dy})
		}
		return r
	}

	rate := (p2.y - p1.y) / (p2.x - p1.x)
	xdis := p2.x - p1.x

	for dx := 1.0; dx < xdis; dx++ {
		r = append(r, Point{x: p1.x + dx, y: p1.y + rate*dx})
	}

	return r
}

func getInput(in string) map[Point]bool {
	res := make(map[Point]bool)
	for y, line := range strings.Split(in, "\n") {
		for x, v := range strings.Split(line, "") {
			if v == "#" {
				res[Point{x: float64(x), y: float64(y)}] = true
			}
		}
	}
	return res
}

func getMaxDetectedAsteroid(in string) (point Point, max int) {
	asteroids := getInput(in)

	for s, _ := range asteroids {
		var count int

		for d, _ := range asteroids {
			var blocked bool

			if s == d {
				continue
			}

			for _, p := range getSigntPoints(s, d) {
				blocked = blocked || asteroids[p]
			}

			if !blocked {
				count++
			}
		}

		if count > max {
			max = count
			point.x = s.x
			point.y = s.y
		}
	}
	return point, max
}

func getLaserCyclePonits(laser Point, left map[Point]bool) []Point {
	var dels []Point

	for del, _ := range left {
		var blocked bool

		for _, mid := range getSigntPoints(laser, del) {
			blocked = blocked || left[mid]
		}

		if !blocked {
			dels = append(dels, del)
		}
	}

	return sortByClockwise(laser, dels)
}

func getArgtagent(center Point, p Point) (argtagent float64) {
	dx := math.Abs(p.x - center.x)
	dy := math.Abs(p.y - center.y)

	if p.x == center.x && p.y < center.y {
		argtagent = 0
	} else if p.x > center.x && p.y < center.y {
		argtagent = math.Atan(dx / dy)
	} else if p.x > center.x && p.y > center.y {
		argtagent = math.Atan(dy/dx) + math.Pi/2
	} else if p.x > center.x && p.y == center.y {
		argtagent = math.Pi / 2
	} else if p.x == center.x && p.y > center.y {
		argtagent = math.Pi
	} else if p.x < center.x && p.y > center.y {
		argtagent = math.Atan(dx/dy) + math.Pi
	} else if p.x < center.x && p.y == center.y {
		argtagent = math.Pi * 3 / 2
	} else if p.x < center.x && p.y < center.y {
		argtagent = math.Atan(dy/dx) + math.Pi*3/2
	}

	return argtagent
}

func sortByClockwise(center Point, points []Point) []Point {
	sort.Slice(points, func(i, j int) (r bool) {
		return getArgtagent(center, points[i]) < getArgtagent(center, points[j])
	})

	return points
}

func getVaporizedPoint(in string, turn int) (p Point) {
	var remain []Point
	left := getInput(in)
	laser, _ := getMaxDetectedAsteroid(in)
	delete(left, laser)

	for turn > 0 {
		dels := getLaserCyclePonits(laser, left)
		if turn < len(dels) || turn == 1 {
			remain = make([]Point, len(dels))
			copy(remain, dels)
			break
		} else {
			turn -= len(dels)

			for _, p := range dels {
				delete(left, p)
			}
		}
	}

	return remain[turn-1]
}

func main() {
	// Part 1
	_, count := getMaxDetectedAsteroid(input)
	fmt.Println(count)
	// Part 2
	p := getVaporizedPoint(input, 200)
	fmt.Println(p.x*100 + p.y)
}

const input string = `###..#########.#####.
.####.#####..####.#.#
.###.#.#.#####.##..##
##.####.#.###########
###...#.####.#.#.####
#.##..###.########...
#.#######.##.#######.
.#..#.#..###...####.#
#######.##.##.###..##
#.#......#....#.#.#..
######.###.#.#.##...#
####.#...#.#######.#.
.######.#####.#######
##.##.##.#####.##.#.#
###.#######..##.#....
###.##.##..##.#####.#
##.########.#.#.#####
.##....##..###.#...#.
#..#.####.######..###
..#.####.############
..##...###..#########`
