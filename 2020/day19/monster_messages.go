/*
  code reference: https://github.com/mnml/aoc/blob/master/2020/19/2.go
*/

package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

var rules = map[string]string{}
var messages = []string{}

func parseInput(filename string) {
	dat, _ := ioutil.ReadFile(filename)
	input := strings.Split(strings.TrimSpace(string(dat)), "\n\n")
	for _, line := range strings.Split(strings.TrimSpace(input[0]), "\n") {
		rule := strings.Split(line, ": ")
		rules[rule[0]] = rule[1]
	}
	for _, line := range strings.Split(strings.TrimSpace(input[1]), "\n") {
		messages = append(messages, line)
	}
}

func regex(ruleID string) (re string) {
	if rules[ruleID][0] == '"' {
		return rules[ruleID][1 : len(rules[ruleID])-1]
	}
	for _, s := range strings.Split(rules[ruleID], " | ") {
		re += "|"
		for _, id := range strings.Fields(s) {
			re += regex(id)
		}
	}
	return "(?:" + re[1:] + ")"
}

func countMessagesMatched(ruleID string, part int) (r int) {
	for _, msg := range messages {
		if part == 2 {
			rules["8"] = `"` + regex("42") + `+"`
			rules["11"] = ""
			for i := 1; i <= 10; i++ {
				rules["11"] += fmt.Sprintf("|%s{%d}%s{%d}", regex("42"), i, regex("31"), i)
			}
			rules["11"] = `"(?:` + rules["11"][1:] + `)"`
		}
		if regexp.MustCompile("(?m)^" + regex(ruleID) + "$").MatchString(msg) {
			r++
		}
	}
	return r
}

func main() {
	parseInput("input")
	count1 := countMessagesMatched("0", 1)
	count2 := countMessagesMatched("0", 2)
	fmt.Println(count1, count2)
}
