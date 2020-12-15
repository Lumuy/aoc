/*
	slow version, it take less than 10 seconds
*/

package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

func parseInput(filename string) (r []int) {
	dat, _ := ioutil.ReadFile(filename)
	for _, s := range strings.Split(string(dat), ",") {
		n, _ := strconv.Atoi(strings.TrimSpace(s))
		r = append(r, n)
	}
	return r
}

func getSpokenNumber(input []int, turns int) int {
	var idx, last int
	m := make(map[int][]int)
	for _, n := range input {
		idx++
		m[n] = []int{idx}
		last = n
	}
	for idx < turns {
		idx++
		var val int

		if len(m[last]) > 1 {
			val = m[last][1] - m[last][0]
		}

		switch len(m[val]) {
		case 0:
			m[val] = []int{idx}
		case 1:
			m[val] = []int{m[val][0], idx}
		case 2:
			m[val] = []int{m[val][1], idx}
		}
		last = val
	}

	return last
}

func main() {
	now := time.Now()
	defer func() {
		fmt.Println("take times", time.Since(now))
	}()

	input := parseInput("input")
	numb1 := getSpokenNumber(input, 2020)
	numb2 := getSpokenNumber(input, 30000000)
	fmt.Println(numb1, numb2)
}
