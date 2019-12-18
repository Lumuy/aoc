package main

import (
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
	collectedKeys := make([]string, 0)

	for {
		nextUnit := Unit{}
		entrance := getEntrance(area)

		reachable_keys := reachableKeys(area, entrance)
		if len(reachable_keys) == 0 {
			break
		}
		fmt.Println("------reachable keys------->")
		for _, u := range reachable_keys {
			fmt.Printf("%s, ", u.name)
		}
		fmt.Println()

		// key and door pair
		for _, v := range reachable_keys {
			nextUnit = v
		}

		// clear matched door
		for p, s := range area {
			if s == strings.ToUpper(nextUnit.name) {
				area[p] = "."
			}
		}
		// move from entrance to key
		area[entrance] = "."
		entrance = nextUnit.point
		area[entrance] = "@"
		steps += nextUnit.steps
		collectedKeys = append(collectedKeys, nextUnit.name)
	}

	fmt.Println("--------Collected order------> ", collectedKeys)

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
