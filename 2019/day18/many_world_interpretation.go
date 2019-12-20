package main

import (
	intcode "aoc/2019/intcode_program"
	"fmt"
	"io/ioutil"
	"strings"
)

const allDoors string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const allKeys string = "abcdefghijklmnopqrstuvwxyz"

type Point struct {
	x, y int
}

type Unit struct {
	point Point
	name  string
	steps int
}

type Movement struct {
	gotKeys     []string
	curEntrance Point
	curSteps    int
}

func getCurrentArea(in string, m Movement) map[Point]string {
	area := make(map[Point]string)
	in = strings.Trim(in, "\n")

	openedDoors := make([]string, 0)
	for _, key := range m.gotKeys {
		openedDoors = append(openedDoors, strings.ToUpper(key))
	}

	y := 0
	for _, line := range strings.Split(in, "\n") {
		for x, c := range strings.Split(line, "") {
			p := Point{x, y}
			if intcode.SliceContains(m.gotKeys, c) ||
				intcode.SliceContains(openedDoors, c) ||
				c == "@" {
				area[p] = "."
			} else {
				area[p] = c
			}
		}
		y++
	}
	area[m.curEntrance] = "@"

	return area
}

func increaseMovement(in string, m Movement) ([]Movement, bool) {
	ms := make([]Movement, 0)
	area := getCurrentArea(in, m)

	curReachableKeys := reachableKeys(area, m.curEntrance)
	if len(curReachableKeys) == 0 {
		return ms, false
	}
	// fmt.Println("----reachable keys------>", curReachableKeys)

	for _, u := range curReachableKeys {
		steps := m.curSteps + u.steps
		keys := append(m.gotKeys, u.name)

		ms = append(ms, Movement{keys, u.point, steps})
	}

	return ms, true
}

func getArea(input string) map[Point]string {
	area := make(map[Point]string)
	input = strings.Trim(input, "\n")

	y := 0
	for _, line := range strings.Split(input, "\n") {
		for x, c := range strings.Split(line, "") {
			area[Point{x, y}] = c
		}
		y++
	}

	return area
}

func pointsContains(points []Point, point Point) bool {
	for _, p := range points {
		if p == point {
			return true
		}
	}
	return false
}

func stepPoint(area map[Point]string, cur Point) ([]Point, map[Point]string) {
	points := make([]Point, 0)
	keys := make(map[Point]string)

	for _, p := range []Point{{0, -1}, {1, 0}, {0, 1}, {-1, 0}} {
		np := Point{p.x + cur.x, p.y + cur.y}
		if area[np] == "." {
			points = append(points, np)
		} else if strings.Contains(allKeys+allDoors, area[np]) {
			keys[np] = area[np]
		}
	}

	return points, keys
}

func reachableKeys(area map[Point]string, entrance Point) (keys []Unit) {
	points := []Point{entrance}
	walked := []Point{entrance}

	for steps := 1; len(points) > 0; steps++ {
		newPoints := make([]Point, 0)

		for _, p := range points {
			ps, ks := stepPoint(area, p)

			for _, p := range ps {
				if !pointsContains(walked, p) {
					walked = append(walked, p)
					newPoints = append(newPoints, p)
				}
			}
			for p, k := range ks {
				if strings.Contains(allKeys, k) {
					keys = append(keys, Unit{p, k, steps})
				}
			}
		}
		points = newPoints
	}

	return keys
}

func getEntrance(area map[Point]string) (entrance Point) {
	for p, c := range area {
		if c == "@" {
			entrance = p
			break
		}
	}

	return entrance
}

func shortestPathSteps(input string) (steps int) {
	area := getArea(input)
	entrance := getEntrance(area)
	movements := []Movement{{[]string{}, entrance, 0}}

	for {
		goon := true
		newMovements := make([]Movement, 0)

		for _, m := range movements {
			ms, done := increaseMovement(input, m)
			newMovements = append(newMovements, ms...)
			goon = goon && done
		}

		if !goon {
			break
		}

		movements = newMovements

		fmt.Println("-----increase movements size---collected keys number---> ", len(movements), len(movements[0].gotKeys))
	}

	for i, m := range movements {
		if i == 0 || steps > m.curSteps {
			steps = m.curSteps
		}
	}

	return steps
}

func main() {
	{
		// Part 1
		dat, _ := ioutil.ReadFile("input")
		input := string(dat)
		steps := shortestPathSteps(input)
		fmt.Println(steps)
	}
}
