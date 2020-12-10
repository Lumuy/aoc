package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"sort"
	"strconv"
	"strings"
)

func parseJolts(filename string) (r []int) {
	dat, _ := ioutil.ReadFile(filename)
	for _, line := range strings.Split(strings.TrimSpace(string(dat)), "\n") {
		n, _ := strconv.Atoi(line)
		r = append(r, n)
	}
	sort.Ints(r)
	return append(r, r[len(r)-1]+3) // device built-int joltage adapter
}

func getMultipliedNum(jolts []int) int {
	var outlet, j1, j3 int // charing outlet, 1-jolt, 3-jolt
	m := make(map[int]int)

	for i := 0; i < len(jolts); i++ {
		m[jolts[i]] = jolts[i] - outlet
		outlet = jolts[i]
	}
	for _, v := range m {
		switch v {
		case 1:
			j1++
		case 3:
			j3++
		}
	}

	return j1 * j3
}

func getNumberOfValidOptions(arr []int) int {
	n := len(arr)
	if arr[n-1]-arr[0]+1 != n {
		panic("Tt's not continous array which logic not fit")
	}
	return int(math.Pow(2, float64(n-1))) - (n-1)/3
}

func countArrangements(jolts []int) int {
	var outlet int
	var scope [][]int
	for i := 0; i < len(jolts); {
		var ele []int
		for jolts[i]-outlet < 3 {
			ele = append(ele, jolts[i])
			outlet = jolts[i]
			i++
		}
		if jolts[i]-outlet == 3 {
			outlet = jolts[i]
			i++
		}
		if len(ele) > 0 {
			scope = append(scope, ele)
		}
	}
	res := 1
	for _, arr := range scope {
		res *= getNumberOfValidOptions(arr)
	}
	return res
}

func main() {
	jolts := parseJolts("input")
	numbr := getMultipliedNum(jolts)
	count := countArrangements(jolts)
	fmt.Println(numbr, count)
}
