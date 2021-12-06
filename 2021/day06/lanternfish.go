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

func count(arr []int, days int) int {
	for d := 0; d < days; d++ {
		var tmp []int
		for _, n := range arr {
			if n == 0 {
				tmp = append(tmp, 6)
				tmp = append(tmp, 8)
			} else {
				tmp = append(tmp, n - 1)
			}
		}
		arr = tmp
	}

	return len(arr)
}

func main() {
	input := parseInput("input")
	p1 := count(input, 80)

	fmt.Println(p1)
}
