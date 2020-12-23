package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"sort"
	"strings"
)

type Food struct {
	is, as []string
}

func parseInput(filename string) (r []Food) {
	dat, _ := ioutil.ReadFile(filename)
	for _, line := range strings.Split(strings.TrimSpace(string(dat)), "\n") {
		var is, as []string
		re := regexp.MustCompile(`^([a-z\s]+) \(contains ([a-z\,\s]+)\)$`)
		ar := re.FindStringSubmatch(line)
		for _, s := range strings.Fields(ar[1]) {
			is = append(is, s)
		}
		for _, s := range strings.Split(ar[2], ", ") {
			as = append(as, s)
		}
		r = append(r, Food{is, as})
	}
	return r
}

func intersection(la, ra []string) (ia []string) {
	m := map[string]bool{}
	for _, s := range la {
		m[s] = true
	}
	for _, s := range ra {
		if m[s] {
			ia = append(ia, s)
		}
	}
	return ia
}

func remove(arr []string, ele string) (r []string) {
	for _, str := range arr {
		if str != ele {
			r = append(r, str)
		}
	}
	return r
}

func find(m map[string][]string) map[string]string {
	r := map[string]string{}

	for len(m) != 0 {
		n := map[string][]string{}
		for a, iarr := range m {
			if len(iarr) == 1 {
				r[iarr[0]] = a
			} else {
				n[a] = iarr
			}
		}
		for a, iarr := range n {
			for i := range r {
				iarr = remove(iarr, i)
			}
			n[a] = iarr
		}
		m = n
	}

	return r
}

func count(foods []Food) (int, string) {
	m := map[string][][]string{}
	n := map[string][]string{}
	r := 0

	for _, food := range foods {
		for _, allergen := range food.as {
			m[allergen] = append(m[allergen], food.is)
		}
	}

	for k, arr := range m {
		ca := arr[0]
		for _, ia := range arr[1:] {
			ca = intersection(ca, ia)
		}
		n[k] = ca
	}

	mr := find(n)
	for _, food := range foods {
		for _, i := range food.is {
			if _, ok := mr[i]; !ok {
				r++
			}
		}
	}

	var is, as []string
	mm := map[string]string{}
	for i, a := range mr {
		mm[a] = i
		as = append(as, a)
	}
	sort.Strings(as)
	for _, a := range as {
		is = append(is, mm[a])
	}

	return r, strings.Join(is, ",")
}

func main() {
	foods := parseInput("input")
	fmt.Println(count(foods))
}
