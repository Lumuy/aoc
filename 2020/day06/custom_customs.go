package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func parseAnswers(filename string) (r [][][]rune) {
	dat, _ := ioutil.ReadFile(filename)
	for _, line := range strings.Split(string(dat), "\n\n") {
		var group [][]rune
		for _, ele := range strings.Split(strings.TrimSpace(line), "\n") {
			group = append(group, []rune(ele))
		}
		r = append(r, group)
	}
	return r
}

func intersection(l, r []rune) (s []rune) {
	m := make(map[rune]bool)
	n := make(map[rune]bool)
	for _, e := range l {
		m[e] = true
	}
	for _, e := range r {
		n[e] = true
	}
	for _, k := range l {
		if m[k] && n[k] {
			s = append(s, k)
		}
	}
	return s
}

func getSumOfCounts(answers [][][]rune) (r, s int) {
	for _, group := range answers {
		m := make(map[rune]int)
		n := group[0]

		for _, p := range group {
			n = intersection(n, p)
			for _, q := range p {
				m[q] += 1
			}
		}

		r += len(m)
		s += len(n)
	}
	return r, s
}

func main() {
	answers := parseAnswers("input")
	fmt.Println(getSumOfCounts(answers))
}
