package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

const tree = '#'

func parseInput(filename string) (r []string) {
	dat, _ := ioutil.ReadFile(filename)
	for _, line := range strings.Split(string(dat), "\n") {
		if line != "" {
			r = append(r, line)
		}
	}
	return r
}

func countTrees(area []string, dx, dy int) (r int) {
	var x int
	size := len(area[0])

	for y := 0; y < len(area)-1; y += dy {
		x += dx
		line := []rune(area[y+dy])
		if line[x%size] == tree {
			r++
		}
	}
	return r
}

func getMultiply(area []string) int {
	res := 1
	for _, pos := range [][]int{{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2}} {
		res *= countTrees(area, pos[0], pos[1])
	}
	return res
}

func main() {
	area := parseInput("input")
	fmt.Println(countTrees(area, 3, 1))
	fmt.Println(getMultiply(area))
}
