package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Point struct {
	x, y, z int
}

type Moon struct {
	pos, vel Point
}

func getInput(in string) (moons []Moon) {
	for _, line := range strings.Split(in, "\n") {
		xyz := regexp.MustCompile(`[-]?\d+`).FindAllString(line, -1)
		x, _ := strconv.Atoi(xyz[0])
		y, _ := strconv.Atoi(xyz[1])
		z, _ := strconv.Atoi(xyz[2])

		moons = append(moons, Moon{pos: Point{x, y, z}})
	}
	return moons
}

func gravity(x1, x2 int) int {
	if x1 < x2 {
		return 1
	} else if x1 > x2 {
		return -1
	} else {
		return 0
	}
}

func move(moons []Moon) []Moon {
	for i, _ := range moons {
		for j, _ := range moons {
			if i == j {
				continue
			}
			moons[i].vel.x += gravity(moons[i].pos.x, moons[j].pos.x)
			moons[i].vel.y += gravity(moons[i].pos.y, moons[j].pos.y)
			moons[i].vel.z += gravity(moons[i].pos.z, moons[j].pos.z)
		}
	}

	for i, _ := range moons {
		moons[i].pos.x += moons[i].vel.x
		moons[i].pos.y += moons[i].vel.y
		moons[i].pos.z += moons[i].vel.z
	}

	return moons
}

func intAbs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

/*
	mode
	0 - x Axix euqal
	1 - y Axix equal
	2 - y Axix equal
*/
func equalMoonsAxis(l, r []Moon, mode int) bool {
	for i, _ := range l {
		switch mode {
		case 0:
			if !(l[i].pos.x == r[i].pos.x && r[i].vel.x == 0) {
				return false
			}
		case 1:
			if !(l[i].pos.y == r[i].pos.y && r[i].vel.y == 0) {
				return false
			}
		case 2:
			if !(l[i].pos.z == r[i].pos.z && r[i].vel.z == 0) {
				return false
			}
		}
	}
	return true
}

func leastCommonMultiple(x1, x2 int) int {
	max_common := 1
	max_number := x1

	if x2 > x1 {
		max_number = x2
	}

	for i := 1; i <= max_number; i++ {
		if x1%i == 0 && x2%i == 0 {
			max_common = i
		}
	}

	return x1 * x2 / max_common
}

func process(in string, count int) (total, steps int) {
	var xc, yc, zc int
	var rx, ry, rz bool
	before := getInput(in)
	after := make([]Moon, len(before))
	copy(after, before)

	for i := 0; ; i++ {
		after = move(after)

		if i+1 == count {
			for _, m := range after {
				pot := intAbs(m.pos.x) + intAbs(m.pos.y) + intAbs(m.pos.z)
				kin := intAbs(m.vel.x) + intAbs(m.vel.y) + intAbs(m.vel.z)

				total += pot * kin
			}
		}

		if !rx && equalMoonsAxis(before, after, 0) {
			xc = i + 1
			rx = true
		}
		if !ry && equalMoonsAxis(before, after, 1) {
			yc = i + 1
			ry = true
		}
		if !rz && equalMoonsAxis(before, after, 2) {
			zc = i + 1
			rz = true
		}

		if xc != 0 && yc != 0 && zc != 0 {
			break
		}
	}

	steps = leastCommonMultiple(xc, yc)
	steps = leastCommonMultiple(zc, steps)

	return total, steps
}

func main() {
	total, steps := process(input, 100)
	fmt.Println(total)
	fmt.Println(steps)
}

const input string = `<x=5, y=-1, z=5>
<x=0, y=-14, z=2>
<x=16, y=4, z=0>
<x=18, y=1, z=16>`
