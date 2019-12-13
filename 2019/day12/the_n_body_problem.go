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

type MoonSet struct {
	moons []Moon
}

func getInput(in string) MoonSet {
	var moons []Moon
	for _, line := range strings.Split(in, "\n") {
		xyz := regexp.MustCompile(`[-]?\d+`).FindAllString(line, -1)
		x, _ := strconv.Atoi(xyz[0])
		y, _ := strconv.Atoi(xyz[1])
		z, _ := strconv.Atoi(xyz[2])

		moons = append(moons, Moon{pos: Point{x, y, z}})
	}
	return MoonSet{moons}
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

func (m *MoonSet) move() {
	for i, _ := range m.moons {
		for j, _ := range m.moons {
			if i == j {
				continue
			}
			m.moons[i].vel.x += gravity(m.moons[i].pos.x, m.moons[j].pos.x)
			m.moons[i].vel.y += gravity(m.moons[i].pos.y, m.moons[j].pos.y)
			m.moons[i].vel.z += gravity(m.moons[i].pos.z, m.moons[j].pos.z)
		}
	}

	for i, _ := range m.moons {
		m.moons[i].pos.x += m.moons[i].vel.x
		m.moons[i].pos.y += m.moons[i].vel.y
		m.moons[i].pos.z += m.moons[i].vel.z
	}
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
func equalMoonsAxis(l, r *MoonSet, mode int) bool {
	for i, _ := range l.moons {
		switch mode {
		case 0:
			if !(l.moons[i].pos.x == r.moons[i].pos.x && l.moons[i].vel.x == r.moons[i].vel.x) {
				return false
			}
		case 1:
			if !(l.moons[i].pos.y == r.moons[i].pos.y && l.moons[i].vel.y == r.moons[i].vel.y) {
				return false
			}
		case 2:
			if !(l.moons[i].pos.z == r.moons[i].pos.z && l.moons[i].vel.z == r.moons[i].vel.z) {
				return false
			}
		}
	}
	return true
}

func gcd(min, max int) int {
	if min > max {
		min, max = max, min
	}
	if min == 0 {
		return max
	} else {
		return gcd(min, max%min)
	}
}

func lcm(x1, x2 int) int {
	return x1 * x2 / gcd(x1, x2)
}

func process(in string, count int) (total, steps int) {
	var xc, yc, zc int
	var rx, ry, rz bool

	data := getInput(in)
	before := &data
	moons := make([]Moon, len(before.moons))
	copy(moons[:], before.moons)
	after := &MoonSet{moons}

	for i := 0; ; i++ {
		after.move()

		if i+1 == count {
			for _, m := range after.moons {
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

	steps = lcm(xc, yc)
	steps = lcm(zc, steps)

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
