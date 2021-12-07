package main

import (
  "fmt"
  "io/ioutil"
  "strconv"
  "strings"
)

func parseInput(filename string) (r []int) {
  dat, _ := ioutil.ReadFile(filename)
  for _, s := range strings.Split(strings.TrimSpace(string(dat)), ",") {
    n, _ := strconv.Atoi(s)
    r = append(r, n)
  }

  return r
}

func max(arr []int) (r int) {
  for i, n := range arr {
    if i == 0 {
      r = n
    } else {
      if n > r {
        r = n
      }
    }
  }

  return r
}

func abs(x int) int {
  if x < 0 {
    return -x
  }
  return x
}

func burnedFuels(x1, x2 int, constant bool) (r int) {
  if constant {
    r =  abs(x1 - x2)
  } else {
    for x := 1; x <= abs(x1 - x2); x++ {
      r += x
    }
  }

  return r
}

func countCheapestFuel(crabs []int, constant bool) (prev, r int) {
  for x := 0; x <= max(crabs); x++ {
    var totalFuels int
    for _, n := range crabs {
      totalFuels += burnedFuels(n, x, constant)
    }
    if x == 0 {
      prev = totalFuels
    } else {
      if totalFuels < prev {
        prev = totalFuels
        r = x
      }
    }
  }

  return prev, r
}

func main() {
  input := parseInput("input")
  p1, _ := countCheapestFuel(input, true)
  p2, _ := countCheapestFuel(input, false)

  fmt.Println(p1)
  fmt.Println(p2)
}
