package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type notes struct {
	rules         map[string][][]int
	yourTicket    []int
	nearbyTickets [][]int
}

func parseInt(str string) int {
	n, _ := strconv.Atoi(strings.TrimSpace(str))
	return n
}

func parseTicket(line string) (r []int) {
	for _, str := range strings.Split(strings.TrimSpace(line), ",") {
		r = append(r, parseInt(str))
	}
	return r
}

func parseNotes(filename string) notes {
	dat, _ := ioutil.ReadFile(filename)
	context := strings.Split(string(dat), "\n\n")

	rules := make(map[string][][]int)
	for _, line := range strings.Split(strings.TrimSpace(context[0]), "\n") {
		re := regexp.MustCompile(`([a-z\s]+): (\d+)-(\d+) or (\d+)-(\d+)`)
		arr := re.FindStringSubmatch(line)
		rules[arr[1]] = [][]int{{parseInt(arr[2]), parseInt(arr[3])}, {parseInt(arr[4]), parseInt(arr[5])}}
	}

	cxt := strings.Split(strings.TrimSpace(context[1]), "\n")[1]
	yourTicket := parseTicket(cxt)

	var nearbyTickets [][]int
	for _, line := range strings.Split(strings.TrimSpace(context[2]), "\n")[1:] {
		nearbyTickets = append(nearbyTickets, parseTicket(line))
	}

	return notes{rules: rules, yourTicket: yourTicket, nearbyTickets: nearbyTickets}
}

func validateTicket(rules map[string][][]int, ticket []int) (r int) {
	for _, n := range ticket {
		var valid bool
		for _, v := range rules {
			if (n >= v[0][0] && n <= v[0][1]) || (n >= v[1][0] && n <= v[1][1]) {
				valid = true
				break
			}
		}
		if !valid {
			r += n
		}
	}
	return r
}

func getTicketScanningErrorRate(nts notes) (validTickets [][]int, r int) {
	for _, ticket := range nts.nearbyTickets {
		errRate := validateTicket(nts.rules, ticket)
		if errRate > 0 {
			r += errRate
		} else {
			validTickets = append(validTickets, ticket)
		}
	}
	return validTickets, r
}

func getValidFields(rules map[string][][]int, val int) (r []string) {
	for key, rule := range rules {
		if (val >= rule[0][0] && val <= rule[0][1]) || (val >= rule[1][0] && val <= rule[1][1]) {
			r = append(r, key)
		}
	}
	return r
}

func intersection(left, right []string) (res []string) {
	m := make(map[string]bool)
	for _, e := range left {
		m[e] = true
	}
	for _, k := range right {
		if _, ok := m[k]; ok {
			res = append(res, k)
		}
	}
	return res
}

func getField(fields, seen []string) (res string) {
	m := make(map[string]bool)
	for _, field := range seen {
		m[field] = true
	}
	for _, field := range fields {
		if _, ok := m[field]; !ok {
			res = field
		}
	}
	return res
}

func getMultipliedNum(nts notes, tickets [][]int) int {
	m, n := make(map[int][]string), make(map[string]int)

	for i := 0; i < len(nts.rules); i++ {
		var fields []string
		for key := range nts.rules {
			fields = append(fields, key)
		}

		for j := 0; j < len(tickets); j++ {
			fields = intersection(fields, getValidFields(nts.rules, tickets[j][i]))
		}
		m[i] = fields
	}

	length, seen := 1, []string{}
	for length < len(m)+1 {
		for idx, fields := range m {
			if len(fields) == length {
				add := getField(fields, seen)
				seen = append(seen, add)
				n[add] = idx
				break
			}
		}
		length++
	}

	r, re := 1, regexp.MustCompile(`^departure`)
	for field, idx := range n {
		if re.MatchString(field) {
			r *= nts.yourTicket[idx]
		}
	}

	return r
}

func main() {
	nts := parseNotes("input")
	validTickets, errRate := getTicketScanningErrorRate(nts)
	num := getMultipliedNum(nts, validTickets)
	fmt.Println(errRate, num)
}
