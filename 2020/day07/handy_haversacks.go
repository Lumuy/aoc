package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

// Bags ..
type Bags struct {
	color string
	num   int
}

func parseRules(filename string) map[Bags][]Bags {
	res := make(map[Bags][]Bags)
	dat, _ := ioutil.ReadFile(filename)
	reK := regexp.MustCompile(`([a-z\s]+) bags contain`)
	reV := regexp.MustCompile(`(\d) ([a-z\s]+) bag[s]?`)
	for _, line := range strings.Split(strings.TrimSpace(string(dat)), "\n") {
		key := Bags{num: 1, color: reK.FindAllStringSubmatch(line, -1)[0][1]}
		var val []Bags
		for _, ele := range reV.FindAllStringSubmatch(line, -1) {
			num, _ := strconv.Atoi(ele[1])
			val = append(val, Bags{num: num, color: ele[2]})
		}
		res[key] = val
	}

	return res
}

func contain(arr []string, ele string) bool {
	for _, e := range arr {
		if e == ele {
			return true
		}
	}
	return false
}

func removeDup(arr []string) (res []string) {
	m := make(map[string]bool)
	for _, e := range arr {
		m[e] = true
	}
	for k := range m {
		res = append(res, k)
	}
	return res
}

func getBagColors(rules map[Bags][]Bags, goals []string) (colors []string) {
	for key, arr := range rules {
		for _, bag := range arr {
			if contain(goals, bag.color) {
				colors = append(colors, key.color)
			}
		}
	}

	return colors
}

func countBagColors(rules map[Bags][]Bags, des string) int {
	m := make(map[string]bool)
	colors := getBagColors(rules, []string{des})
	for len(colors) != 0 {
		for _, e := range colors {
			m[e] = true
		}
		var tmp []string
		for _, color := range colors {
			res := getBagColors(rules, []string{color})
			tmp = append(tmp, res...)
		}
		colors = removeDup(tmp)
	}
	return len(m)
}

func countIndividualBags(rules map[Bags][]Bags, des string) (r int) {
	for _, bags := range rules[Bags{num: 1, color: des}] {
		r += bags.num * (1 + countIndividualBags(rules, bags.color))
	}

	return r
}

func main() {
	rules := parseRules("input")
	fmt.Println(countBagColors(rules, "shiny gold"))
	fmt.Println(countIndividualBags(rules, "shiny gold"))
}
