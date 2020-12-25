/*
	code reference: https://github.com/mnml/aoc/blob/master/2020/20/2.go
*/

package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"regexp"
	"strconv"
	"strings"
)

type Tile []string

type Point struct {
	x, y int
}

func (t Tile) Col(n int) (r string) {
	for _, s := range t {
		r += string(s[n])
	}
	return r
}

func (t Tile) Edges() []string {
	return []string{t[0], t.Col(len(t[0]) - 1), t[len(t)-1], t.Col(0)}
}

func (t Tile) rotateAndFlip() []Tile {
	r := make([]Tile, 8)
	for i := 0; i < 8; i += 2 {
		for _, s := range t {
			r[i] = append(r[i], reverse(s))
		}
		for j := range t {
			r[i+1] = append(r[i+1], reverse(t.Col(j)))
		}
		t = r[i+1]
	}
	return r
}

func reverse(s string) (rs string) {
	for _, r := range s {
		rs = string(r) + rs
	}
	return rs
}

func parseInput(filename string) map[int]Tile {
	m := map[int]Tile{}
	dat, _ := ioutil.ReadFile(filename)
	for _, cxt := range strings.Split(strings.TrimSpace(string(dat)), "\n\n") {
		lines := strings.Split(strings.TrimSpace(cxt), "\n")
		idstr := regexp.MustCompile(`\d+`).FindString(lines[0])
		id, _ := strconv.Atoi(idstr)

		tile := Tile{}
		for _, line := range lines[1:] {
			tile = append(tile, line)
		}
		m[id] = tile
	}
	return m
}

func run(tiles map[int]Tile) (int, int) {
	part1, part2 := 1, 0
	monster := []string{"..................#.", "#....##....##....###", ".#..#..#..#..#..#..."}
	counts := map[string]int{}
	for _, t := range tiles {
		for _, e := range t.Edges() {
			counts[e]++
			counts[reverse(e)]++
		}
	}

	imageSize, tileSize := int(math.Sqrt(float64(len(tiles)))), 10
	image := make(Tile, imageSize*(tileSize-2))
	order := map[Point]Tile{}

	for y := 0; y < imageSize; y++ {
		for x := 0; x < imageSize; x++ {
		findTile:
			for id, t := range tiles {
				for _, pt := range t.rotateAndFlip() {
					yDirection := y == 0 && counts[pt[0]] == 1 || y != 0 && pt[0] == order[Point{x, y - 1}][tileSize-1]
					xDirection := x == 0 && counts[pt.Col(0)] == 1 || x != 0 && pt.Col(0) == order[Point{x - 1, y}].Col(tileSize-1)
					if xDirection && yDirection {
						if (y == 0 || y == imageSize-1) && (x == 0 || x == imageSize-1) {
							part1 *= id
						}

						for i := 0; i < tileSize-2; i++ {
							image[(tileSize-2)*y+i] += pt[i+1][1 : tileSize-1]
						}

						order[Point{x, y}] = pt
						delete(tiles, id)
						break findTile
					}
				}
			}
		}
	}

	var monsterNum int
	for _, r := range image.rotateAndFlip() {
		for y := 0; y < len(r)-len(monster); y++ {
		findMonsters:
			for x := 0; x < len(r[0])-len(monster[0]); x++ {
				for i, s := range monster {
					if match, _ := regexp.MatchString(s, r[y+i][x:x+len(s)]); !match {
						continue findMonsters
					}
				}
				monsterNum++
			}
		}
	}
	part2 = strings.Count(strings.Join(image, ""), "#") - monsterNum*strings.Count(strings.Join(monster, ""), "#")

	return part1, part2
}

func main() {
	tiles := parseInput("input")
	fmt.Println(run(tiles))
}
