package main

import "testing"

func Test(t *testing.T) {
  input := parseInput("test_input")
  p1, _ := countCheapestFuel(input, true)
  p2, _ := countCheapestFuel(input, false)

  if p1 != 37 {
    t.Errorf("Part one should be 37, but be %d", p1)
  }

  if p2 != 168 {
    t.Errorf("Part two should be 168, but be %d", p2)
  }
}
