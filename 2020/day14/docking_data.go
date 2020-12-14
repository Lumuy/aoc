package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

const size = 36

type action struct {
	mem, val int
}

type instruction struct {
	mask    string
	actions []action
}

func parseProgram(filename string) (r []instruction) {
	dat, _ := ioutil.ReadFile(filename)

	var ins instruction
	for _, line := range strings.Split(strings.TrimSpace(string(dat)), "\n") {
		if regexp.MustCompile(`^mask.*`).MatchString(line) {
			if len(ins.actions) != 0 {
				r = append(r, ins)
			}
			re := regexp.MustCompile(`^mask = ([10X]+)`)
			ins = instruction{mask: re.FindStringSubmatch(line)[1]}
		} else {
			arr := regexp.MustCompile(`^mem\[(\d+)\] = (\d+)`).FindStringSubmatch(line)
			m, _ := strconv.Atoi(arr[1])
			v, _ := strconv.Atoi(arr[2])
			ins.actions = append(ins.actions, action{mem: m, val: v})
		}
	}
	r = append(r, ins)

	return r
}

func getFloatingAddresses(arr []rune) (r []string) {
	var depth int
	for _, v := range arr {
		if v == 'X' {
			depth++
		}
	}

	kinds := []string{"1", "0"}
	for i := 1; i < depth; i++ {
		var tmp []string
		for _, str := range kinds {
			tmp = append(tmp, str+"0")
			tmp = append(tmp, str+"1")
		}
		kinds = tmp
	}

	for _, kind := range kinds {
		var index int
		arrMem, arrKind := []rune(string(arr)), []rune(kind)
		for idx, val := range arrMem {
			if val == 'X' {
				arrMem[idx] = arrKind[index]
				index++
			}
		}
		r = append(r, string(arrMem))
	}

	return r
}

func computeMems(mask string, mem int) (r []int) {
	arrMask, arrMem := []rune(mask), []rune(fmt.Sprintf("%036b", mem))
	for i := 0; i < size; i++ {
		if arrMask[i] == 'X' || arrMask[i] == '1' {
			arrMem[i] = arrMask[i]
		}
	}
	for _, str := range getFloatingAddresses(arrMem) {
		n, _ := strconv.ParseInt(str, 2, 64)
		r = append(r, int(n))
	}
	return r
}

func computeVal(mask string, val int) int {
	arrMask, arrVal := []rune(mask), []rune(fmt.Sprintf("%036b", val))
	for i := 0; i < size; i++ {
		if arrMask[i] != 'X' {
			arrVal[i] = arrMask[i]
		}
	}
	r, _ := strconv.ParseInt(string(arrVal), 2, 64)

	return int(r)
}

func getMemorySum(program []instruction, part int) (r int) {
	m := make(map[int]int)

	for _, ins := range program {
		for _, action := range ins.actions {
			if part == 1 {
				m[action.mem] = computeVal(ins.mask, action.val)
			}
			if part == 2 {
				for _, mem := range computeMems(ins.mask, action.mem) {
					m[mem] = action.val
				}
			}
		}
	}
	for _, val := range m {
		r += val
	}
	return r
}

func main() {
	program := parseProgram("input")
	sum1 := getMemorySum(program, 1)
	sum2 := getMemorySum(program, 2)
	fmt.Println(sum1, sum2)
}
