package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Layer [5][5]bool

/*
    layers         origin
                     |
    <----------------|----------------->
 Index               |
                     0
                  -1 0 +1     <------ one minutes
               -2 -1 0 +1 +2  <------ two minutes
                 ..........
 Each minute layers spread both the outer and inner diretion.
*/
func bugsCount(lines []string) (count int) {
	var origin Layer
	for y, line := range lines {
		for x, char := range strings.Split(line, "") {
			if char == "#" {
				origin[y][x] = true
			}
		}
	}

	curLayers := make(map[int]Layer)
	curLayers[0] = origin
	min, max := 0, 0

	for minute := 0; minute < 200; minute++ {
		nextLayers := make(map[int]Layer)

		for index := min - 1; index <= max+1; index++ {
			var nextLayer Layer
			inner, outer, cur := curLayers[index-1], curLayers[index+1], curLayers[index]

			for y := 0; y < 5; y++ {
				for x := 0; x < 5; x++ {
					neighbors := 0

					if x == 2 && y == 2 {
						continue
					}
					if y == 0 && outer[2][1] {
						neighbors++
					}
					if y == 4 && outer[2][3] {
						neighbors++
					}
					if x == 0 && outer[1][2] {
						neighbors++
					}
					if x == 4 && outer[3][2] {
						neighbors++
					}

					if x == 1 && y == 2 {
						for i := 0; i < 5; i++ {
							if inner[0][i] {
								neighbors++
							}
						}
					}

					if x == 3 && y == 2 {
						for i := 0; i < 5; i++ {
							if inner[4][i] {
								neighbors++
							}
						}
					}

					if x == 2 && y == 1 {
						for i := 0; i < 5; i++ {
							if inner[i][0] {
								neighbors++
							}
						}
					}

					if x == 2 && y == 3 {
						for i := 0; i < 5; i++ {
							if inner[i][4] {
								neighbors++
							}
						}
					}

					if x > 0 && cur[x-1][y] {
						neighbors++
					}
					if x < 4 && cur[x+1][y] {
						neighbors++
					}
					if y > 0 && cur[x][y-1] {
						neighbors++
					}
					if y < 4 && cur[x][y+1] {
						neighbors++
					}

					nextLayer[x][y] = (cur[x][y] && neighbors == 1) || (!cur[x][y] && neighbors >= 1 && neighbors <= 2)
				}
			}

			nextLayers[index] = nextLayer
		}

		curLayers = nextLayers
		min--
		max++
	}

	for _, layer := range curLayers {
		for y := 0; y < 5; y++ {
			for x := 0; x < 5; x++ {
				if layer[x][y] {
					count++
				}
			}
		}
	}

	return count
}

func biodiversityRating(lines []string) uint32 {
	var state uint32

	for y, line := range lines {
		for x, char := range strings.Split(line, "") {
			if char == "#" {
				state |= (1 << uint32(5*y+x))
			}
		}
	}

	seen := make(map[uint32]bool)
	for !seen[state] {
		seen[state] = true

		var next uint32

		for x := 0; x < 5; x++ {
			for y := 0; y < 5; y++ {
				neighbors := 0

				if x > 0 && state&(1<<uint32(5*y+x-1)) != 0 {
					neighbors++
				}
				if x < 4 && state&(1<<uint32(5*y+x+1)) != 0 {
					neighbors++
				}
				if y > 0 && state&(1<<uint32(5*(y-1)+x)) != 0 {
					neighbors++
				}
				if y < 4 && state&(1<<uint32(5*(y+1)+x)) != 0 {
					neighbors++
				}
				var bit uint32 = 1 << uint32(5*y+x)
				if ((state&bit != 0) && neighbors == 1) || ((state&bit == 0) && neighbors >= 1 && neighbors <= 2) {
					next |= bit
				}
			}
		}

		state = next
	}

	return state
}

func main() {
	lines := readLines("input")

	// Part 1
	fmt.Println(biodiversityRating(lines))

	// Part 2
	fmt.Println(bugsCount(lines))
}

func readLines(filename string) []string {
	dat, err := ioutil.ReadFile(filename)
	check(err)
	return strings.Split(string(dat), "\n")
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
