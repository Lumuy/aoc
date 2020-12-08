package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

// Instruction ...
type Instruction struct {
	operation string
	argument  int
}

func getInstructions(filename string) (r []Instruction) {
	dat, _ := ioutil.ReadFile(filename)
	for _, line := range strings.Split(strings.TrimSpace(string(dat)), "\n") {
		arr := strings.Split(line, " ")
		n, _ := strconv.Atoi(arr[1])
		r = append(r, Instruction{operation: arr[0], argument: n})
	}
	return r
}

func contain(arr []int, ele int) bool {
	for _, val := range arr {
		if val == ele {
			return true
		}
	}
	return false
}

func getAccumulatorValue(ins []Instruction) (int, bool) {
	var idx, val int
	var handled []int
	terminated := true

	for {
		if contain(handled, idx) {
			terminated = false
			break
		}
		if idx > len(ins)-1 {
			break
		}
		handled = append(handled, idx)
		switch ins[idx].operation {
		case "acc":
			val += ins[idx].argument
			idx++
		case "jmp":
			idx += ins[idx].argument
		case "nop":
			idx++
		}
	}

	return val, terminated
}

func getProgramValue(ins []Instruction) (r int) {
	for idx, ele := range ins {
		fixed := make([]Instruction, len(ins))
		copy(fixed, ins)
		switch ele.operation {
		case "jmp":
			fixed[idx].operation = "nop"
		case "nop":
			fixed[idx].operation = "jmp"
		default:
			continue
		}
		value, terminated := getAccumulatorValue(fixed)
		if terminated {
			r = value
			break
		}
	}

	return r
}

func main() {
	ins := getInstructions("input")
	val, _ := getAccumulatorValue(ins)
	cnt := getProgramValue(ins)
	fmt.Println(val, cnt)
}
