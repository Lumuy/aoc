package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Policy struct {
	min, max   int
	ch, passwd string
}

func parseInput(filename string) (r []Policy) {
	dat, _ := ioutil.ReadFile(filename)

	for _, line := range strings.Split(string(dat), "\n") {
		if line != "" {
			arr := strings.Split(line, " ")
			values := strings.Split(arr[0], "-")
			min, _ := strconv.Atoi(values[0])
			max, _ := strconv.Atoi(values[1])
			ch := strings.Replace(arr[1], ":", "", 1)
			r = append(r, Policy{min: min, max: max, ch: ch, passwd: arr[2]})
		}
	}

	return r
}

func countValidPasswords(passwords []Policy) (r int) {
	for _, policy := range passwords {
		var count int
		com := []rune(policy.ch)[0]
		for _, ele := range []rune(policy.passwd) {
			if ele == com {
				count++
			}
		}
		if count >= policy.min && count <= policy.max {
			r++
		}
	}
	return r
}

func countValidPasswordsTwo(passwords []Policy) (r int) {
	for _, policy := range passwords {
		com := []rune(policy.ch)[0]
		pas := []rune(policy.passwd)
		if (com == pas[policy.min-1]) != (com == pas[policy.max-1]) {
			r++
		}
	}
	return r
}

func main() {
	passwords := parseInput("input")
	fmt.Println(countValidPasswords(passwords))
	fmt.Println(countValidPasswordsTwo(passwords))
}
