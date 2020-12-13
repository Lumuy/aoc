package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func parseBuses(filename string) (start int, buses []int) {
	dat, _ := ioutil.ReadFile(filename)
	arr := strings.Split(strings.TrimSpace(string(dat)), "\n")
	start, _ = strconv.Atoi(arr[0])
	for _, bus := range strings.Split(arr[1], ",") {
		n, _ := strconv.Atoi(bus)
		buses = append(buses, n)
	}
	return start, buses
}

func getMultipliedNumber(start int, buses []int) (r int) {
	var tval int
	m := make(map[int]int)

	for _, id := range buses {
		if id == 0 {
			continue
		}
		for val := 0; ; val += id {
			if val > start {
				m[val-start] = id
				break
			}
		}
	}
	for val, _ := range m {
		if tval == 0 || tval > val {
			tval = val
		}
	}

	return tval * m[tval]
}

func getEarliestTimestamp(buses []int) (r int) {
	jump := 1
	for idx, bid := range buses {
		if bid == 0 {
			continue
		}
		for (r+idx)%bid != 0 {
			r += jump
		}
		jump *= bid
	}

	return r
}

func main() {
	start, buses := parseBuses("input")
	num := getMultipliedNumber(start, buses)
	tim := getEarliestTimestamp(buses)
	fmt.Println(num, tim)
}
