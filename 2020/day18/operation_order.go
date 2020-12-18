package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func parseInput(filename string) (h []string) {
	dat, _ := ioutil.ReadFile(filename)
	for _, s := range strings.Split(strings.TrimSpace(string(dat)), "\n") {
		h = append(h, s)
	}
	return h
}

func parseInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func eval(s string, part int) (r int) {
	switch part {
	case 1:
		a := strings.Fields(s)
		r = parseInt(a[0])
		for i := 1; i < len(a); i += 2 {
			switch a[i] {
			case "*":
				r *= parseInt(a[i+1])
			case "+":
				r += parseInt(a[i+1])
			}
		}
	case 2:
		for strings.Contains(s, "+") {
			a := strings.Fields(s)
			for i, v := range a {
				if v == "+" {
					bs := strings.Join(a[:i-1], " ")
					as := strings.Join(a[i+2:], " ")
					s = fmt.Sprintf("%s %d %s", bs, parseInt(a[i-1])+parseInt(a[i+1]), as)
					break
				}
			}
		}

		r = 1
		for _, v := range strings.Fields(s) {
			if v != "*" {
				r *= parseInt(v)
			}
		}
	}

	return r
}

func count(s string, part int) int {
	for strings.Contains(s, "(") {
		var l, r int
		for i, e := range []rune(s) {
			if e == '(' {
				l = i
			}
			if e == ')' {
				r = i
				break
			}
		}
		s = fmt.Sprintf("%s%d%s", s[0:l], count(s[l+1:r], part), s[r+1:])
	}
	return eval(s, part)
}

func sum(h []string, part int) (r int) {
	for _, s := range h {
		r += count(s, part)
	}
	return r
}

func main() {
	inp := parseInput("input")
	cn1, cn2 := sum(inp, 1), sum(inp, 2)
	fmt.Println(cn1, cn2)
}
