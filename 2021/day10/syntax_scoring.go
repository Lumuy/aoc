package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

func parseInput(filename string) (r []string) {
	dat, _ := ioutil.ReadFile(filename)
	for _, s := range strings.Split(strings.TrimSpace(string(dat)), "\n") {
		r = append(r, s)
	}
	return r
}

func reduce(line string) (int, string) {
	var n int
	var arr []rune
	size := len(line)

	for i := 0; i < size; {
		if i+1 < size {
			cond1 := line[i] == '(' && line[i+1] == ')'
			cond2 := line[i] == '[' && line[i+1] == ']'
			cond3 := line[i] == '{' && line[i+1] == '}'
			cond4 := line[i] == '<' && line[i+1] == '>'

			if cond1 || cond2 || cond3 || cond4 {
				n++
				i += 2
			} else {
				arr = append(arr, rune(line[i]))
				i++
			}
		} else {
			arr = append(arr, rune(line[i]))
			i++
		}
	}

	return n, string(arr)
}

func contains(arr []rune, ele rune) bool {
	for _, r := range arr {
		if r == ele {
			return true
		}
	}
	return false
}

func countScores(lines []string) (score, middle int) {
	var arr []string
	for _, line := range lines {
		n := 1
		for n != 0 {
			n, line = reduce(line)
		}
		arr = append(arr, line)
	}

	var incompletes []string
	for _, line := range arr {
		var isIllegal bool
	findFirst:
		for _, r := range line {
			if !contains([]rune{'(', '[', '{', '<'}, r) {
				isIllegal = true
				switch r {
				case ')':
					score += 3
				case ']':
					score += 57
				case '}':
					score += 1197
				case '>':
					score += 25137
				}

				break findFirst
			}
		}

		if !isIllegal {
			incompletes = append(incompletes, line)
		}
	}

	var scores []int
	for _, str := range incompletes {
		n := 0
		for i := len(str) - 1; i > -1; i-- {
			n = n * 5
			switch rune(str[i]) {
			case '(':
				n += 1
			case '[':
				n += 2
			case '{':
				n += 3
			case '<':
				n += 4
			}
		}
		scores = append(scores, n)
	}
	sort.Ints(scores)
	middle = scores[len(scores)/2]

	return score, middle
}

func main() {
	input := parseInput("input")
	p1, p2 := countScores(input)

	fmt.Println(p1)
	fmt.Println(p2)
}
