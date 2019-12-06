package main

import "testing"

func Test(t *testing.T) {
  number := getTotalOrbits(input_part1)
  count := getMinimumTransfers(input_part2, "YOU", "SAN")
  if number != 42 {
    t.Errorf("Test orbits should be 42, but %d", number)
  }
  if count != 4 {
    t.Errorf("Test transfers should be 4, but %d", count)
  }
}

const input_part1 string = `COM)B
B)C
C)D
D)E
E)F
B)G
G)H
D)I
E)J
J)K
K)L`

const input_part2 string = `COM)B
B)C
C)D
D)E
E)F
B)G
G)H
D)I
E)J
J)K
K)L
K)YOU
I)SAN`
