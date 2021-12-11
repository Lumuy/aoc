package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func parseInput(filename string) (r []int) {
	dat, _ := ioutil.ReadFile(filename)
	for _, s := range strings.Split(strings.TrimSpace(string(dat)), ",") {
		n, _ := strconv.Atoi(s)
		r = append(r, n)
	}
	return r
}

func step(n, days int) (count int) {
	if n < days {
		count += step(6, days-n-1)
		count += step(8, days-n-1)
	} else {
		count = 1
	}

	return count
}

func countElement(arr []int, ele int) (r int) {
	for _, val := range arr {
		if ele == val {
			r++
		}
	}

	return r
}

func count(arr []int, days int) (r int) {
	for n := 1; n < 6; n++ {
		r += countElement(arr, n) * step(n, days)
	}

	return r
}

func main() {
	input := parseInput("input")
	p1 := count(input, 80)
	p2 := count(input, 256)

	fmt.Println(p1)
	fmt.Println(p2)
}
