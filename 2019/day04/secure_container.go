package main

import (
	"fmt"
	"strconv"
	"strings"
)

func convertToSortedArray(y int) [6]int {
	x1, y1 := y/100000, y%100000
	x2, y2 := y1/10000, y1%10000
	x3, y3 := y2/1000, y2%1000
	x4, y4 := y3/100, y3%100
	x5, x6 := y4/10, y4%10

	return [6]int{x1, x2, x3, x4, x5, x6}
}

func isValidPassword(y int) bool {
	x := convertToSortedArray(y)

	if x[5] >= x[4] && x[4] >= x[3] && x[3] >= x[2] && x[2] >= x[1] && x[1] >= x[0] {
		if x[0] == x[1] || x[1] == x[2] || x[2] == x[3] || x[3] == x[4] || x[4] == x[5] {
			return true
		}
	}

	return false
}

func isNotPartOfLargerGroup(y int) bool {
	hash := make(map[int]int)

	for _, x := range convertToSortedArray(y) {
		hash[x] += 1
	}
	for _, v := range hash {
		if v == 2 {
			return true
		}
	}

	return false
}

func countDifferentPasswords(input string) (count []int) {
	arr := strings.Split(input, "-")
	min, _ := strconv.Atoi(arr[0])
	max, _ := strconv.Atoi(arr[1])

	for i := min; i <= max; i++ {
		if isValidPassword(i) {
			count = append(count, i)
		}
	}

	return count
}

func main() {
	arr := countDifferentPasswords(input)
	var count int
	// Part 1
	fmt.Println(len(arr))
	// Part 2
	for _, x := range arr {
		if isNotPartOfLargerGroup(x) {
			count++
		}
	}
	fmt.Println(count)
}

const input string = `172851-675869`
