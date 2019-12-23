package main

import (
	intcode "aoc/2019/intcode_program"
	"fmt"
)

type Point struct {
	X, Y int
}

type QueueItem struct {
	Position Point
	Distance int
	Next     *QueueItem
}

var (
	Wall = 0
	Path = 1
)

var (
	North = int64(1)
	South = int64(2)
	West  = int64(3)
	East  = int64(4)
)

var (
	Up    = Point{0, -1}
	Down  = Point{0, 1}
	Left  = Point{-1, 0}
	Right = Point{1, 0}
)

var (
	directions = map[int64]Point{North: Up, South: Down, West: Left, East: Right}
	reverses   = map[int64]int64{North: South, South: North, West: East, East: West}
)

func main() {
	program := intcode.ParseInput("input")

	input := make(chan int64)
	output := make(chan int64)
	halt := make(chan bool)

	go intcode.Run(program, input, output, halt)

	pos := Point{0, 0}

	area := make(map[Point]int)
	area[pos] = Path

	var oxygenPos Point
	var oxygenDistance int

	{
		var queue []QueueItem
		queue = append(queue, QueueItem{Position: pos, Distance: 0})

		for len(queue) != 0 {
			item := queue[0]
			queue = queue[1:]

			pos = move(pos, item.Position, area, input, output)

			for dr, dp := range directions {
				next, nextDistance := item.Position.Add(dp), item.Distance+1
				if _, ok := area[next]; !ok {
					input <- dr
					switch <-output {
					case 0:
						area[next] = Wall
					case 2:
						if oxygenDistance == 0 {
							oxygenPos = next
							oxygenDistance = nextDistance
						}
						fallthrough
					case 1:
						area[next] = Path
						queue = append(queue, QueueItem{Position: next, Distance: nextDistance})
						// reverse direction for other situation
						input <- reverses[dr]
						<-output
					}
				}
			}
		}
	}

	// Part 1
	fmt.Println(oxygenDistance)

	var minDistance int

	// Fill complete map and record maximum distance
	{
		var queue []QueueItem
		queue = append(queue, QueueItem{Position: oxygenPos, Distance: 0})

		visited := make(map[Point]bool)
		visited[oxygenPos] = true

		for len(queue) != 0 {
			item := queue[0]
			queue = queue[1:]

			minDistance = intcode.MaxInt(minDistance, item.Distance)

			for _, dp := range directions {
				next := item.Position.Add(dp)
				if !visited[next] && area[next] == Path {
					visited[next] = true
					queue = append(queue, QueueItem{Position: next, Distance: item.Distance + 1})
				}
			}
		}
	}

	// Part 2
	fmt.Println(minDistance)
	// printArea(area, oxygenPos)
}

func move(pos, target Point, area map[Point]int, input chan int64, output chan int64) Point {
	var link *QueueItem

	// Find shortest route from target to pos (in reversed order)
	var queue []QueueItem
	queue = append(queue, QueueItem{Position: target, Distance: 0})

	visited := make(map[Point]bool)
	visited[target] = true

	for len(queue) != 0 {
		item := queue[0]
		queue = queue[1:]

		if item.Position == pos {
			link = &item
			break
		}

		for _, dp := range directions {
			next := item.Position.Add(dp)
			if !visited[next] && area[next] == Path {
				visited[next] = true
				queue = append(queue, QueueItem{Position: next, Distance: item.Distance + 1, Next: &item})
			}
		}
	}

	// Follow path backwards from pos to target
	for link.Next != nil {
		for dr, dp := range directions {
			if link.Position.Add(dp) == link.Next.Position {
				input <- dr
				<-output
				break
			}
		}

		link = link.Next
	}

	return target
}

func (p Point) Add(dp Point) Point {
	return Point{p.X + dp.X, p.Y + dp.Y}
}

func (l Point) Min(r Point) Point {
	return Point{
		X: intcode.MinInt(l.X, r.X),
		Y: intcode.MinInt(l.Y, r.Y),
	}
}

func (l Point) Max(r Point) Point {
	return Point{
		X: intcode.MaxInt(l.X, r.X),
		Y: intcode.MaxInt(l.Y, r.Y),
	}
}

func printArea(area map[Point]int, oxygenPos Point) {
	var min, max Point

	for pos := range area {
		min = min.Min(pos)
		max = max.Max(pos)
	}

	res := ""
	for y := min.Y; y <= max.Y; y++ {
		for x := min.X; x <= max.X; x++ {
			pos := Point{x, y}
			value, ok := area[pos]
			if x == 0 && y == 0 {
				res += "S"
			} else if pos == oxygenPos {
				res += "O"
			} else if !ok {
				res += "?"
			} else if value == Path {
				res += " "
			} else {
				res += "#"
			}
		}
		res += "\n"
	}
	fmt.Println(res)
}
