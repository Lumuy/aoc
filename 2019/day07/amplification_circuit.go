package main

import (
	"fmt"
	"strconv"
	"strings"
)

func getInput(input string) (res []int) {
	for _, s := range strings.Split(input, ",") {
		add, _ := strconv.Atoi(s)
		res = append(res, add)
	}

	return res
}

func fetchValue(data []int, pos int, mode int) (int, bool) {
	if mode == 0 {
		return data[pos], true
	} else if mode == 1 {
		return pos, true
	}
	return 0, false
}

func process(data []int, pointer, input int) (bool, []int, int, int) {
	var res []string
	var hang bool
	var keep bool
	var halts bool

	for !halts {
		v := data[pointer]
		a, x := v/10000, v%10000
		b, x := x/1000, x%1000
		c, de := x/100, x%100
		if a != 0 {
			fmt.Println("Error destination mode")
		}

		switch de {
		case 1:
			l, _ := fetchValue(data, data[pointer+1], c)
			r, _ := fetchValue(data, data[pointer+2], b)
			d := data[pointer+3]

			data[d] = l + r
			pointer += 4
		case 2:
			l, _ := fetchValue(data, data[pointer+1], c)
			r, _ := fetchValue(data, data[pointer+2], b)
			d := data[pointer+3]

			data[d] = l * r
			pointer += 4
		case 3:
			if !hang {
				d := data[pointer+1]
				data[d] = input
				pointer += 2
				hang = true
			} else {
				keep = true
				halts = true
			}
		case 4:
			v, _ = fetchValue(data, data[pointer+1], c)
			pointer += 2
			res = append(res, strconv.Itoa(v))
		case 5:
			v1, _ := fetchValue(data, data[pointer+1], c)
			v2, _ := fetchValue(data, data[pointer+2], b)
			if v1 != 0 {
				pointer = v2
			} else {
				pointer += 3
			}
		case 6:
			v1, _ := fetchValue(data, data[pointer+1], c)
			v2, _ := fetchValue(data, data[pointer+2], b)
			if v1 == 0 {
				pointer = v2
			} else {
				pointer += 3
			}
		case 7:
			v1, _ := fetchValue(data, data[pointer+1], c)
			v2, _ := fetchValue(data, data[pointer+2], b)
			v3 := data[pointer+3]
			if v1 < v2 {
				data[v3] = 1
			} else {
				data[v3] = 0
			}
			pointer += 4
		case 8:
			v1, _ := fetchValue(data, data[pointer+1], c)
			v2, _ := fetchValue(data, data[pointer+2], b)
			v3 := data[pointer+3]
			if v1 == v2 {
				data[v3] = 1
			} else {
				data[v3] = 0
			}
			pointer += 4
		case 99:
			halts = true
		default:
			fmt.Println("Unknown opcode: ", a, b, c, de)
			halts = true
		}
	}

	output, _ := strconv.Atoi(strings.Join(res[:], ""))
	return keep, data, pointer, output
}

func addInt(left, right int) int {
	i := 10
	for {
		if right/i == 0 {
			break
		}
		i *= 10
	}
	return left*i + right
}

func getMaxSignal(input string) (max int) {
	for i1 := 0; i1 < 5; i1++ {
		for i2 := 0; i2 < 5; i2++ {
			if i2 == i1 {
				continue
			}
			for i3 := 0; i3 < 5; i3++ {
				if i3 == i1 || i3 == i2 {
					continue
				}
				for i4 := 0; i4 < 5; i4++ {
					if i4 == i1 || i4 == i2 || i4 == i3 {
						continue
					}
					for i5 := 0; i5 < 5; i5++ {
						if i5 == i1 || i5 == i2 || i5 == i3 || i5 == i4 {
							continue
						}

						output := 0
						for _, phase := range []int{i1, i2, i3, i4, i5} {
							_, data, pointer, out1 := process(getInput(input), 0, phase)
							_, _, _, out2 := process(data, pointer, output)
							output = addInt(out1, out2)
						}

						if output > max {
							max = output
						}
					}
				}
			}
		}
	}

	return max
}

func getMaxFeedbackSignal(input string) (max int) {
	for i1 := 5; i1 < 10; i1++ {
		for i2 := 5; i2 < 10; i2++ {
			if i2 == i1 {
				continue
			}
			for i3 := 5; i3 < 10; i3++ {
				if i3 == i2 || i3 == i1 {
					continue
				}
				for i4 := 5; i4 < 10; i4++ {
					if i4 == i3 || i4 == i2 || i4 == i1 {
						continue
					}
					for i5 := 5; i5 < 10; i5++ {
						if i5 == i4 || i5 == i3 || i5 == i2 || i5 == i1 {
							continue
						}
						// TODO
						var in int
						var index int
						var pointers [5]int
						var data [5][]int
						var halts [5]bool

						for di, phase := range []int{i1, i2, i3, i4, i5} {
							_, d, p, _ := process(getInput(input), pointers[di], phase)
							data[di] = d
							pointers[di] = p
						}

						for {
							keep, d, p, ou := process(data[index], pointers[index], in)
							in = ou
							data[index] = d
							pointers[index] = p

							if !keep {
								halts[index] = true
							}
							if halts == [5]bool{true, true, true, true, true} {
								break
							}

							if index == 4 {
								index = 0
							} else {
								index++
							}
						}

						if in > max {
							max = in
						}
					}
				}
			}
		}
	}
	return max
}

func main() {
	// Part 1
	fmt.Println(getMaxSignal(input))
	// Part 2
	fmt.Println(getMaxFeedbackSignal(input))
}

const input string = `3,8,1001,8,10,8,105,1,0,0,21,42,63,76,101,114,195,276,357,438,99999,3,9,101,2,9,9,102,5,9,9,1001,9,3,9,1002,9,5,9,4,9,99,3,9,101,4,9,9,102,5,9,9,1001,9,5,9,102,2,9,9,4,9,99,3,9,1001,9,3,9,1002,9,5,9,4,9,99,3,9,1002,9,2,9,101,5,9,9,102,3,9,9,101,2,9,9,1002,9,3,9,4,9,99,3,9,101,3,9,9,102,2,9,9,4,9,99,3,9,1001,9,2,9,4,9,3,9,102,2,9,9,4,9,3,9,101,2,9,9,4,9,3,9,102,2,9,9,4,9,3,9,101,1,9,9,4,9,3,9,1001,9,2,9,4,9,3,9,1001,9,1,9,4,9,3,9,1001,9,2,9,4,9,3,9,1001,9,2,9,4,9,3,9,1001,9,1,9,4,9,99,3,9,102,2,9,9,4,9,3,9,1001,9,2,9,4,9,3,9,102,2,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,1001,9,1,9,4,9,3,9,102,2,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,102,2,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,102,2,9,9,4,9,99,3,9,102,2,9,9,4,9,3,9,102,2,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,1001,9,2,9,4,9,3,9,1002,9,2,9,4,9,3,9,1001,9,1,9,4,9,3,9,1002,9,2,9,4,9,3,9,1002,9,2,9,4,9,3,9,101,2,9,9,4,9,3,9,1001,9,2,9,4,9,99,3,9,1001,9,1,9,4,9,3,9,101,2,9,9,4,9,3,9,102,2,9,9,4,9,3,9,1001,9,2,9,4,9,3,9,1001,9,1,9,4,9,3,9,102,2,9,9,4,9,3,9,1001,9,2,9,4,9,3,9,1001,9,2,9,4,9,3,9,102,2,9,9,4,9,3,9,1001,9,2,9,4,9,99,3,9,102,2,9,9,4,9,3,9,101,1,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,1002,9,2,9,4,9,3,9,1002,9,2,9,4,9,3,9,101,2,9,9,4,9,3,9,1001,9,2,9,4,9,3,9,101,2,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,101,2,9,9,4,9,99`
