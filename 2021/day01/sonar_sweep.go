package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func parseInput(filename string) (r []int) {
	dat, _ := ioutil.ReadFile(filename)
	for _, line := range strings.Split(strings.TrimSpace(string(dat)), "\n") {
		n, _ := strconv.Atoi(line)
		r = append(r, n)
	}

	return r
}

func three_measurement(x []int) (y []int) {
	for i := 0; i < len(x)-2; i++ {
		y = append(y, x[i]+x[i+1]+x[i+2])
	}

	return y
}

func count(arr []int) (r int) {
	var pre int

	for i, n := range arr {
		if i != 0 && n > pre {
			r++
		}
		pre = n
	}

	return r
}

func main() {
	arr1 := parseInput("input")
	arr2 := three_measurement(arr1)

	fmt.Println(count(arr1))
	fmt.Println(count(arr2))
}
