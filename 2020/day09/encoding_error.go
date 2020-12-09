package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func parseNumbers(filename string) (r []int) {
	dat, _ := ioutil.ReadFile(filename)
	for _, line := range strings.Split(strings.TrimSpace(string(dat)), "\n") {
		n, _ := strconv.Atoi(line)
		r = append(r, n)
	}
	return r
}

func contain(arr []int, ele int) bool {
	for _, val := range arr {
		if ele == val {
			return true
		}
	}
	return false
}

func isFollowRule(preamble []int, num int) bool {
	for _, val := range preamble {
		rest := num - val
		if rest > 0 && val != rest && contain(preamble, rest) {
			return true
		}
	}
	return false
}

func getFirstNumberUnfollowRule(numbers []int, offset int) (r int) {
	for idx := offset; idx < len(numbers); idx++ {
		if !isFollowRule(numbers[idx-offset:idx], numbers[idx]) {
			r = numbers[idx]
			break
		}
	}
	return r
}

func isMatchScope(numbers []int, num int) bool {
	var total int
	for _, val := range numbers {
		total += val
	}
	return total == num
}

func getEncryptionWeakness(numbers []int, num int) int {
	var scope []int
	for idx := 0; idx < len(numbers); idx++ {
		for offset := 1; offset <= len(numbers)-idx; offset++ {
			if offset > 1 && isMatchScope(numbers[idx:idx+offset], num) {
				scope = numbers[idx : idx+offset]
				break
			}
		}
	}
	sort.Ints(scope)
	return scope[0] + scope[len(scope)-1]
}

func main() {
	numbers := parseNumbers("input")
	num := getFirstNumberUnfollowRule(numbers, 25)
	wks := getEncryptionWeakness(numbers, num)
	fmt.Println(num, wks)
}
