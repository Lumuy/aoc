package main

import (
	intcode "aoc/2019/intcode_program"
	"fmt"
	"io/ioutil"
)

type Point struct {
	x, y int
}

func pointPulled(point Point) bool {
	dat, _ := ioutil.ReadFile("inputs")
	inputs := string(dat)

	p := &intcode.Program{Mem: intcode.GetInput(inputs)}
	p = intcode.Process(p, point.x)
	p = intcode.PrepareBeforeRun(p)
	p = intcode.Process(p, point.y)

	return p.Msg[0] == 1
}

func affectedPointsNumber(size int) (num int) {
	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			if pointPulled(Point{x, y}) {
				num++
			}
		}
	}

	return num
}

func closestPoint() Point {
	x, y := 0, 100

	for {
		if !pointPulled(Point{x, y}) {
			x++
		}

		if pointPulled(Point{x, y}) && // Bottom left
			pointPulled(Point{x + 99, y}) && // Bottom right
			pointPulled(Point{x, y - 99}) && // Top right
			pointPulled(Point{x + 99, y - 99}) { // Top right
			return Point{x, y - 99}
		}

		y++
	}
}

func main() {
	{
		// Part 1
		num := affectedPointsNumber(50)
		fmt.Println(num)
	}

	{
		// Part 2
		p := closestPoint()
		fmt.Println(p.x*10000 + p.y)
	}
}
