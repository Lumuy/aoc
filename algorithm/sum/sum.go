// 数组[2,3,5,7,9]，输出任意组合，可以重复选，输出所有和是 13 的组合

package main

import (
	"fmt"
	"sort"
)

var list = []int{2, 3, 5, 7, 9}
var resl = [][]int{}

func process(total int, eles []int) {
	for _, v := range list {
		left := total - v
		keep := []int{v}
		for _, ele := range eles {
			keep = append(keep, ele)
		}
		if left == 0 {
			sort.Ints(keep)
			resl = append(resl, keep)
		} else if left > 0 {
			process(left, keep)
		}
	}
}

func isSameArray(left, right []int) bool {
	if len(left) == len(right) {
		for k, v := range left {
			if v != right[k] {
				return false
			}
		}
	} else {
		return false
	}
	return true
}

func removeDuplicate(in [][]int) (res [][]int) {
	for _, arr := range in {
		add := true
		for _, ele := range res {
			if isSameArray(ele, arr) {
				add = false
			}
		}
		if add {
			res = append(res, arr)
		}
	}
	return res
}

func main() {
	process(13, []int{})
	for _, v := range removeDuplicate(resl) {
		fmt.Println(v)
	}
}
