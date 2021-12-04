package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func parseInput(filename string) (r [][]int) {
	dat, _ := ioutil.ReadFile(filename)
	for _, line := range strings.Split(strings.TrimSpace(string(dat)), "\n") {
		add := []int{}
		for _, n := range line {
			x, _ := strconv.Atoi(string(n))
			add = append(add, x)
		}
		r = append(r, add)
	}

	return r
}

func countSupportRating(lines [][]int) int {
	return countRating(lines, true) * countRating(lines, false)
}

func countRating(lines [][]int, large bool) (r int) {
	size := len(lines[0])
	var res []int

	for i := 0; i < size; i++ {
		var arr []int
		for _, line := range lines {
			arr = append(arr, line[i])
		}

		var x, y int
		count := countArrayElement(arr, 1)
		if count > len(lines)-count {
			x = 1
		} else if count < len(lines)-count {
			x = 0
		} else {
			x = 1
		}
		y = 1 - x
		if !large && len(lines) == 1 {
			y = lines[0][i]
		}

		if large {
			res = append(res, x)
		} else {
			res = append(res, y)
		}

		var tmp [][]int
		for _, line := range lines {
			if (line[i] == x && large) || (!large && line[i] == y) {
				tmp = append(tmp, line)
			}
		}
		lines = tmp
	}

	for i, n := range res {
		r += powInt(2, size-1-i) * n
	}

	return r
}

func countConsumption(lines [][]int) int {
	var gamma, epsilon int
	length := len(lines[0])
	for i := length - 1; i >= 0; i-- {
		var arr []int
		var x, y int
		for _, line := range lines {
			arr = append(arr, line[i])
		}

		count := countArrayElement(arr, 1)
		if count > len(lines)-count {
			x = 1
		} else {
			x = 0
		}
		y = 1 - x

		gamma += powInt(2, length-1-i) * x
		epsilon += powInt(2, length-1-i) * y
	}

	return gamma * epsilon
}

func powInt(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}

func countArrayElement(arr []int, ele int) (r int) {
	for _, n := range arr {
		if ele == n {
			r++
		}
	}

	return r
}

func main() {
	input := parseInput("input")
	p1 := countConsumption(input)
	p2 := countSupportRating(input)

	fmt.Println(p1)
	fmt.Println(p2)
}
