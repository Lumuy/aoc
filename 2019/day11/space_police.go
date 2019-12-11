package main

import (
	"aoc/2019/intcode_program"
	"fmt"
)

type Point struct {
	x, y int
}

type Panel struct {
	location  Point
	direction int
}

/*
	Paint direction
	0 UP
	1 LEFT
	2 DOWN
	3 RIGHT

	turn
	0 - turn left
	1 - turn right
*/
func getNextPanel(panel Panel, turn int) Panel {
	var dx, dy, direction int
	x, y := panel.location.x, panel.location.y

	switch panel.direction {
	case 0:
		if turn == 0 {
			dx = -1
			direction = 3
		} else if turn == 1 {
			dx = 1
			direction = 1
		}
	case 1:
		if turn == 0 {
			dy = 1
			direction = 0
		} else if turn == 1 {
			dy = -1
			direction = 2
		}
	case 2:
		if turn == 0 {
			dx = 1
			direction = 1
		} else if turn == 1 {
			dx = -1
			direction = 3
		}
	case 3:
		if turn == 0 {
			dy = -1
			direction = 2
		} else if turn == 1 {
			dy = 1
			direction = 0
		}
	}

	return Panel{location: Point{x: x + dx, y: y + dy}, direction: direction}
}

func getPaintedPanels(in string, bg int) map[Point]int {
	p := &intcode_program.Program{Mem: intcode_program.GetInput(in)}
	panel := Panel{}
	paints := make(map[Point]int)
	paints[panel.location] = bg
	goon := true

	for goon {
		// reset
		p.Msg = []int{}
		p.Halts = false
		p.Hang = false
		p = intcode_program.Process(p, paints[panel.location])
		goon = p.Hang

		if len(p.Msg) != 2 {
			break
		}

		paints[panel.location] = p.Msg[0]
		panel = getNextPanel(panel, p.Msg[1])
	}

	return paints
}

func printMsg(paints map[Point]int) {
	reverse_y := make(map[Point]int)
	for p, c := range paints {
		reverse_y[Point{p.x, -p.y}] = c
	}

	var minx, maxx, miny, maxy int
	for p, _ := range reverse_y {
		if p.x < minx {
			minx = p.x
		}
		if p.x > maxx {
			maxx = p.x
		}
		if p.y < miny {
			miny = p.y
		}
		if p.y > maxy {
			maxy = p.y
		}
	}

	for y := miny; y <= maxy; y++ {
		for x := minx; x < maxx; x++ {
			color, ok := reverse_y[Point{x, y}]
			if ok && color == 1 {
				fmt.Printf("*")
			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Println()
	}
}

func main() {
	// Part 1
	m1 := getPaintedPanels(input, 0)
	fmt.Println(len(m1))
	// Part 2
	m2 := getPaintedPanels(input, 1)
	printMsg(m2)
}

const input string = `3,8,1005,8,305,1106,0,11,0,0,0,104,1,104,0,3,8,1002,8,-1,10,101,1,10,10,4,10,1008,8,0,10,4,10,1002,8,1,29,3,8,102,-1,8,10,1001,10,1,10,4,10,108,1,8,10,4,10,1002,8,1,50,1,104,20,10,1,1102,6,10,1006,0,13,3,8,102,-1,8,10,101,1,10,10,4,10,108,1,8,10,4,10,102,1,8,83,1,1102,0,10,1006,0,96,2,1004,19,10,3,8,1002,8,-1,10,101,1,10,10,4,10,108,0,8,10,4,10,101,0,8,116,3,8,1002,8,-1,10,1001,10,1,10,4,10,108,1,8,10,4,10,102,1,8,138,1006,0,60,1,1008,12,10,3,8,102,-1,8,10,101,1,10,10,4,10,1008,8,0,10,4,10,102,1,8,168,1006,0,14,1006,0,28,3,8,1002,8,-1,10,1001,10,1,10,4,10,108,0,8,10,4,10,101,0,8,195,2,1005,9,10,1006,0,29,3,8,1002,8,-1,10,101,1,10,10,4,10,108,1,8,10,4,10,1002,8,1,224,2,1009,8,10,1,3,5,10,3,8,1002,8,-1,10,101,1,10,10,4,10,108,1,8,10,4,10,102,1,8,254,3,8,102,-1,8,10,1001,10,1,10,4,10,1008,8,0,10,4,10,1002,8,1,277,1,1003,18,10,1,1104,1,10,101,1,9,9,1007,9,957,10,1005,10,15,99,109,627,104,0,104,1,21101,0,666681062292,1,21102,322,1,0,1105,1,426,21101,847073883028,0,1,21102,333,1,0,1105,1,426,3,10,104,0,104,1,3,10,104,0,104,0,3,10,104,0,104,1,3,10,104,0,104,1,3,10,104,0,104,0,3,10,104,0,104,1,21101,0,179356855319,1,21102,1,380,0,1105,1,426,21102,1,179356998696,1,21102,1,391,0,1105,1,426,3,10,104,0,104,0,3,10,104,0,104,0,21101,0,988669698816,1,21101,0,414,0,1106,0,426,21102,1,868494500628,1,21102,425,1,0,1106,0,426,99,109,2,21202,-1,1,1,21102,1,40,2,21102,457,1,3,21102,1,447,0,1105,1,490,109,-2,2105,1,0,0,1,0,0,1,109,2,3,10,204,-1,1001,452,453,468,4,0,1001,452,1,452,108,4,452,10,1006,10,484,1102,0,1,452,109,-2,2105,1,0,0,109,4,1201,-1,0,489,1207,-3,0,10,1006,10,507,21102,0,1,-3,22101,0,-3,1,21202,-2,1,2,21101,1,0,3,21102,1,526,0,1106,0,531,109,-4,2105,1,0,109,5,1207,-3,1,10,1006,10,554,2207,-4,-2,10,1006,10,554,22101,0,-4,-4,1106,0,622,21201,-4,0,1,21201,-3,-1,2,21202,-2,2,3,21102,573,1,0,1106,0,531,21202,1,1,-4,21101,1,0,-1,2207,-4,-2,10,1006,10,592,21102,1,0,-1,22202,-2,-1,-2,2107,0,-3,10,1006,10,614,22101,0,-1,1,21102,614,1,0,105,1,489,21202,-2,-1,-2,22201,-4,-2,-4,109,-5,2105,1,0`
