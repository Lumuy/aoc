package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

const sum = 2020

func getMultiplyNumber(arr []int) int {
	for i := 0; i < len(arr)-2; i++ {
		for j := i + 1; j < len(arr)-1; j++ {
			if arr[i]+arr[j] == sum {
				return arr[i] * arr[j]
			}
		}
	}

	return 0
}

func getMultiplyNumberThree(arr []int) int {
	for i := 0; i < len(arr)-3; i++ {
		for j := i + 1; j < len(arr)-2; j++ {
			for k := j + 1; k < len(arr)-1; k++ {
				if arr[i]+arr[j]+arr[k] == sum {
					return arr[i] * arr[j] * arr[k]
				}
			}
		}
	}

	return 0
}

func parseIntArray(filename string) (res []int) {
	dat, err := ioutil.ReadFile(filename)
	check(err)

	for _, ele := range strings.Split(string(dat), "\n") {
		if ele != "" {
			i, _ := strconv.Atoi(ele)
			res = append(res, i)
		}
	}

	return res
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	inputs := parseIntArray("input")
	// Part one
	fmt.Println(getMultiplyNumber(inputs))

	// Part two
	fmt.Println(getMultiplyNumberThree(inputs))
}
