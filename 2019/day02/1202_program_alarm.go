package main

import (
	"fmt"
	"strconv"
	"strings"
)

func getInputData(input string) (res []int) {
	arr := strings.Split(input, ",")

	for _, v := range arr {
		add, _ := strconv.Atoi(v)
		res = append(res, add)
	}

	return res
}

func validInstruction(arr []int, max int) bool {
	var r int
	for i, v := range arr {
		if i == 0 || r < v {
			r = v
		}
	}

	return max > r
}

func runProgram(arr []int, max int) (int, bool) {
	for i := 0; ; i += 4 {
		code := arr[i : i+4]

		if !validInstruction(code, max) {
			return arr[0], false
		}

		if code[0] == 1 {
			arr[code[3]] = arr[code[1]] + arr[code[2]]
		} else if code[0] == 2 {
			arr[code[3]] = arr[code[1]] * arr[code[2]]
		} else if code[0] == 99 {
			break
		} else {
			fmt.Println("Unknow opcode!")
			return arr[0], false
		}
	}

	return arr[0], true
}

func main() {
	// Part 1
	data := getInputData(input)
	max := len(data)
	arr := make([]int, max)
	copy(arr, data)

	// Part 1
	arr[1] = 12
	arr[2] = 2
	res, success := runProgram(arr, max)
	if success {
		fmt.Println(res)
	}

	// Part 2
	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			arr := make([]int, max)
			copy(arr, data)
			arr[1] = noun
			arr[2] = verb
			res, success := runProgram(arr, max)

			if success && res == 19690720 {
				fmt.Println(100*noun + verb)
				break
			} else {
				continue
			}
		}
	}
}

const input string = `1,0,0,3,1,1,2,3,1,3,4,3,1,5,0,3,2,13,1,19,1,19,6,23,1,23,6,27,1,13,27,31,2,13,31,35,1,5,35,39,2,39,13,43,1,10,43,47,2,13,47,51,1,6,51,55,2,55,13,59,1,59,10,63,1,63,10,67,2,10,67,71,1,6,71,75,1,10,75,79,1,79,9,83,2,83,6,87,2,87,9,91,1,5,91,95,1,6,95,99,1,99,9,103,2,10,103,107,1,107,6,111,2,9,111,115,1,5,115,119,1,10,119,123,1,2,123,127,1,127,6,0,99,2,14,0,0`
