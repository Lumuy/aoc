package main

import (
  "fmt"
  "io/ioutil"
  "strconv"
  "strings"
  "math"
)

func parseInput(filename string) (r [][]int) {
	dat, _ := ioutil.ReadFile(filename)
	for _, line := range strings.Split(strings.TrimSpace(string(dat)), "\n") {
    add := []int{}
    for _, n := range line {
      x, _ := strconv.Atoi(string(n))
      add = append(add, x)
    }
    r = append(r, add)
  }

  return r
}

func countConsumption(lines [][]int) int {
  var gamma, epsilon int
  length := len(lines[0])
  for i := length - 1; i >= 0; i-- {
    var arr []int
    var x, y int
    for _, line := range lines {
      arr = append(arr, line[i])
    }

    count := countArrayElement(arr, 1)
    if count > len(lines) - count {
      x = 1
    } else {
      x = 0
    }
    y = 1 - x

    gamma += powInt(2, length - 1 - i) * x
    epsilon += powInt(2, length - 1 - i) * y
  }

  return gamma * epsilon
}

func powInt(x, y int) int {
    return int(math.Pow(float64(x), float64(y)))
}

func countArrayElement(arr []int, ele int) (r int) {
  for _, n := range arr {
    if ele == n {
      r++
    }
  }

  return r
}

func main() {
  input := parseInput("input")
  p1 := countConsumption(input)

  fmt.Println(p1)
}
