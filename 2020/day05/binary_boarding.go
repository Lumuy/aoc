package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

func getBoardingPasses(filename string) (r []string) {
	dat, _ := ioutil.ReadFile(filename)
	for _, line := range strings.Split(string(dat), "\n") {
		if line != "" {
			r = append(r, line)
		}
	}
	return r
}

func parsePassId(pass string) int {
	var r_min, c_min int
	r_max := 127
	c_max := 7
	arr := []rune(pass)

	for _, ele := range arr[:7] {
		switch ele {
		case 'F':
			r_max = (r_max-r_min+1)/2 - 1 + r_min
		case 'B':
			r_min += (r_max - r_min + 1) / 2
		}
	}

	for _, ele := range arr[7:] {
		switch ele {
		case 'R':
			c_min += (c_max - c_min + 1) / 2
		case 'L':
			c_max = (c_max-c_min+1)/2 - 1 + c_min
		}
	}

	return r_min*8 + c_min
}

func main() {
	var ids []int
	for _, pass := range getBoardingPasses("input") {
		ids = append(ids, parsePassId(pass))
	}
	sort.Ints(ids)
	fmt.Println(ids[len(ids)-1])

	for i := 0; i < len(ids)-1; i++ {
		if ids[i+1] != ids[i]+1 {
			fmt.Println(ids[i] + 1)
		}
	}
}
