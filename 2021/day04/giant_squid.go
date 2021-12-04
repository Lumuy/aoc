package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type Matrix struct {
	data [5][5]int
	mark [5][5]int
}

func parseInput(filename string) (arr []int, matrixs []Matrix) {
	dat, _ := ioutil.ReadFile(filename)
	re := regexp.MustCompile(`\d+`)
	for index, line := range strings.Split(strings.TrimSpace(string(dat)), "\n\n") {
		if index == 0 {
			for _, s := range strings.Split(line, ",") {
				n, _ := strconv.Atoi(s)
				arr = append(arr, n)
			}
		} else {
			var add, mark [5][5]int
			for j, str := range strings.Split(strings.TrimSpace(line), "\n") {
				var ele [5]int
				for i, s := range re.FindAllString(str, -1) {
					n, _ := strconv.Atoi(s)
					ele[i] = n
				}
				add[j] = ele
			}
			for i := 0; i < 5; i++ {
				for j := 0; j < 5; j++ {
					mark[i][j] = -1
				}
			}
			matrixs = append(matrixs, Matrix{data: add, mark: mark})
		}
	}

	return arr, matrixs
}

func countFinalScore(arr []int, matrixs []Matrix) (p1, p2 int) {
	wins := make(map[int]int)

	for _, n := range arr {
		for idx, matrix := range matrixs {
			for i, x := range matrix.data {
				for j, y := range x {
					if y == n {
						matrixs[idx].mark[i][j] = n
					}
				}
			}
		}

		for i, m := range matrixs {
			end, sum := check(m)
			if end && p1 == 0 {
				p1 = sum * n
			}
			if end && wins[i] == 0 {
				wins[i] = sum * n
				p2 = sum * n
			}
		}
	}

	return p1, p2
}

func check(m Matrix) (bool, int) {
	for i := 0; i < 5; i++ {
		if !contains(m.mark[i], -1) {
			return true, sum(m)
		}

		var row [5]int
		for j := 0; j < 5; j++ {
			row[j] = m.mark[j][i]
		}
		if !contains(row, -1) {
			return true, sum(m)
		}
	}

	return false, 0
}

func sum(m Matrix) (r int) {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if m.mark[i][j] == -1 {
				r += m.data[i][j]
			}
		}
	}
	return r
}

func contains(arr [5]int, ele int) bool {
	for _, n := range arr {
		if n == ele {
			return true
		}
	}
	return false
}

func main() {
	arr, matrix := parseInput("input")
	p1, p2 := countFinalScore(arr, matrix)

	fmt.Println(p1)
	fmt.Println(p2)
}
