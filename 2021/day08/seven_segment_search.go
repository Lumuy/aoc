package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

var digitRelationship = map[string]int{
	"abcefg":  0,
	"cf":      1,
	"acdeg":   2,
	"acdfg":   3,
	"bcdf":    4,
	"abdfg":   5,
	"abdefg":  6,
	"acf":     7,
	"abcdefg": 8,
	"abcdfg":  9,
}

type Entry struct {
	signals, outputs []string
}

func parseInput(filename string) (r []Entry) {
	re := regexp.MustCompile(`[a-g]+`)
	dat, _ := ioutil.ReadFile(filename)
	for _, line := range strings.Split(strings.TrimSpace(string(dat)), "\n") {
		var add Entry
		parts := strings.Split(line, "|")
		add.signals = re.FindAllString(parts[0], -1)
		add.outputs = re.FindAllString(parts[1], -1)
		r = append(r, add)
	}

	return r
}

func uniqSignalLen() (r []int) {
	var arr []int
	for k, _ := range digitRelationship {
		arr = append(arr, len(k))
	}

	m := make(map[int]int)
	for _, n := range arr {
		m[n] = m[n] + 1
	}

	for k, v := range m {
		if v == 1 {
			r = append(r, k)
		}
	}

	return r
}

func contains(arr []int, ele int) bool {
	for _, n := range arr {
		if n == ele {
			return true
		}
	}
	return false
}

func countUniqDigits(entries []Entry) (r int) {
	uniqs := uniqSignalLen()
	for _, entry := range entries {
		for _, output := range entry.outputs {
			if contains(uniqs, len(output)) {
				r++
			}
		}
	}

	return r
}

func sumOutputs(entries []Entry) (r int) {
	return r
}

func main() {
	input := parseInput("input")
	p1 := countUniqDigits(input)
	p2 := sumOutputs(input)

	fmt.Println(p1)
	fmt.Println(p2)
}
