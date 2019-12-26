package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strings"
)

type Point struct {
	X, Y int
}

/*
	Keys number is less than 32
	so a 32 bit unsigned integer to resprent collected keys

	1 ---> collected
	0 ---> uncollected

	n ----bits------ 31 30 29 .... 2 1 0

	3 keys required  <===> 000...111 binary
*/
type Connection struct {
	NextNode     *Node
	RequiredKeys uint32
	Distance     int
}

type Node struct {
	Key         byte
	Position    Point
	Connections []Connection
}

type Item struct {
	Position     Point
	RequiredKeys uint32
	Distance     int
}

type Area map[Point]byte

var originArea = make(map[Point]byte)

func main() {
	input := readFile("input")
	// Part 1
	n := shortestPathSteps(input)
	fmt.Println(n)
}

func shortestPathSteps(input string) int {
	area := parseInput(input)
	for k, v := range area {
		originArea[k] = v
	}
	return process(area)
}

func process(area Area) int {
	keysNumber := 0
	nodes := make(map[byte]*Node)

	for p, char := range area {
		if isKey(char) {
			keysNumber++
		}
		if isEntrance(char) || isKey(char) {
			nodes[char] = &Node{Key: char, Position: p}
		}
	}

	for _, node := range nodes {
		var spreads []Item
		spreads = append(spreads, Item{node.Position, 0, 0})

		seen := make(map[Point]bool)
		seen[node.Position] = true

		// Spread on up, down, left and right directions
		for len(spreads) != 0 {
			current := spreads[0]
			spreads = spreads[1:]

			for _, nextPos := range current.Position.Adjacents() {
				if seen[nextPos] {
					continue
				}

				seen[nextPos] = true
				nextVal := area[nextPos]

				if nextVal == '.' || isEntrance(nextVal) {
					spreads = append(spreads, Item{nextPos, current.RequiredKeys, current.Distance + 1})
				} else if isDoor(nextVal) {
					spreads = append(spreads, Item{nextPos, current.RequiredKeys | bitFromDoor(nextVal), current.Distance + 1})
				} else if isKey(nextVal) {
					spreads = append(spreads, Item{nextPos, current.RequiredKeys | bitFromKey(nextVal), current.Distance + 1})

					node.Connections = append(node.Connections, Connection{
						NextNode:     nodes[nextVal],
						RequiredKeys: current.RequiredKeys,
						Distance:     current.Distance + 1,
					})
				}
			}
		}
	}

	// Find start positions
	var entrance *Node
	for _, node := range nodes {
		if isEntrance(node.Key) {
			entrance = node
		}
	}

	var allKeys uint32 = (1 << keysNumber) - 1
	cs := make([]string, 0)
	return find(entrance, 0, allKeys, 0, math.MaxInt32, cs)
}

func reachableKeys(key byte, collected uint32) (keys uint32) {
	var curPos Point
	for p, char := range originArea {
		if char == key {
			curPos = p
		}
	}

	var spreads []Point
	spreads = append(spreads, curPos)

	seen := make(map[Point]bool)
	seen[curPos] = true

	for len(spreads) != 0 {
		current := spreads[0]
		spreads = spreads[1:]

		for _, p := range current.Adjacents() {
			if seen[p] {
				continue
			}
			seen[p] = true
			val := originArea[p]
			if isDoor(val) {
				continue
			} else if isKey(val) {
				keys = keys | bitFromKey(val)
				continue
			} else if isEntrance(val) || isWall(val) {
				spreads = append(spreads, p)
			}
		}
	}

	return keys
}

func find(node *Node, collectedKeys, allKeys uint32, curDistance, bestDistance int, cs []string) int {
	if keysCollected(collectedKeys, allKeys) {
		return curDistance
	}

	// fmt.Println("collected keys------> ", cs)
	reachable := reachableKeys(node.Key, collectedKeys)

	for _, conn := range node.Connections {
		if !keysCollected(collectedKeys, conn.RequiredKeys) {
			continue
		}

		if keysCollected(collectedKeys, bitFromKey(conn.NextNode.Key)) {
			continue
		}

		// fmt.Printf("%b \n", reachableKeys(node.Key, collectedKeys))
		// fmt.Println(!keysCollected(reachableKeys(node.Key, collectedKeys), bitFromKey(conn.NextNode.Key)))
		// fmt.Println("--------------------------->", cs)
		// fmt.Printf("%b \n", reachableKeys(node.Key, collectedKeys))
		// fmt.Printf("%b %s \n", bitFromKey(conn.NextNode.Key), string(conn.NextNode.Key))
		// fmt.Println(!keysCollected(reachableKeys(node.Key, collectedKeys), bitFromKey(conn.NextNode.Key)))
		if !keysCollected(reachable, bitFromKey(conn.NextNode.Key)) {
			continue
		}

		newDistance := curDistance + conn.Distance
		if newDistance > bestDistance {
			continue
		}

		nextCollectedKeys := collectedKeys | bitFromKey(conn.NextNode.Key)
		newCs := append(cs, string(conn.NextNode.Key))
		distance := find(conn.NextNode, nextCollectedKeys, allKeys, newDistance, bestDistance, newCs)
		bestDistance = minInt(bestDistance, distance)
	}

	return bestDistance
}

func (p Point) Add(dp Point) Point {
	return Point{p.X + dp.X, p.Y + dp.Y}
}

func (p Point) Adjacents() (r []Point) {
	for _, dp := range []Point{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
		r = append(r, p.Add(dp))
	}
	return r
}

func readFile(filename string) string {
	dat, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return strings.TrimSpace(string(dat))
}

func parseInput(input string) Area {
	area := make(map[Point]byte)
	for y, line := range strings.Split(input, "\n") {
		for x, v := range []byte(line) {
			area[Point{x, y}] = v
		}
	}
	return area
}

func isEntrance(char byte) bool {
	return char == '@'
}

func isKey(char byte) bool {
	return char >= 'a' && char <= 'z'
}

func isWall(char byte) bool {
	return char == '#'
}

func isDoor(char byte) bool {
	return char >= 'A' && char <= 'Z'
}

func bitFromKey(key byte) uint32 {
	return 1 << (key - 'a')
}

func bitFromDoor(door byte) uint32 {
	return 1 << (door - 'A')
}

func keysCollected(collected, allKeys uint32) bool {
	return collected&allKeys == allKeys
}

func minInt(l, r int) int {
	if l < r {
		return l
	}
	return r
}
