package main

import (
	"fmt"
)

func isEqualArray(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func contains(arr [][]int, sub []int) bool {
	for _, ele := range arr {
		if isEqualArray(ele, sub) {
			return true
		}
	}
	return false
}

func combination(in []int) (ou [][]int) {
	if len(in) == 1 {
		return [][]int{in}
	}
	for i, v := range in {
		var reduce []int
		reduce = append(reduce, in[0:i]...)
		reduce = append(reduce, in[i+1:len(in)]...)

		for _, arr := range combination(reduce) {
			for j := 0; j < len(arr); j++ {
				var add []int
				add = append(add, arr[0:j]...)
				add = append(add, v)
				add = append(add, arr[j:len(arr)]...)
				if !contains(ou, add) {
					ou = append(ou, add)
				}
			}
		}
	}
	return ou
}

func main() {
	arr := []int{0, 1, 2, 3, 4}
	for _, v := range combination(arr) {
		fmt.Println(v)
	}
	fmt.Println(len(combination(arr)))
}
