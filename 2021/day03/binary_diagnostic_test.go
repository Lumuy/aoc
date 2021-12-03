package main

import "testing"

func Test(t *testing.T) {
  input := parseInput("test_input")
  p1 := countConsumption(input)

  if p1 != 198 {
    t.Errorf("Part one should be 198, but be %d", p1)
  }
}

