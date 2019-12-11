package intcode_program

import (
	"fmt"
	"strconv"
	"strings"
)

type Program struct {
	Pointer, Rbase int
	Halts, Hang    bool
	Mem            map[int]int
	Msg            []int
}

func GetInput(input string) map[int]int {
	res := make(map[int]int)
	for i, v := range strings.Split(input, ",") {
		n, _ := strconv.Atoi(v)
		res[i] = n
	}
	return res
}

func parseParam(p *Program, val, mode int) (res int) {
	switch mode {
	case 0:
		res = p.Mem[val]
	case 1:
		res = val
	case 2:
		res = p.Mem[p.Rbase+val]
	default:
		fmt.Println("Unknown mode!")
	}
	return res
}

func parseAddr(p *Program, val, mode int) (addr int) {
	switch mode {
	case 0:
		addr = val
	case 2:
		addr = p.Rbase + val
	default:
		fmt.Println("Error mode")
	}
	return addr
}

func Process(p *Program, in int) *Program {
	for !p.Halts {
		v := p.Mem[p.Pointer]
		m3, x := v/10000, v%10000
		m2, x := x/1000, x%1000
		m1, opcode := x/100, x%100

		v1 := parseParam(p, p.Mem[p.Pointer+1], m1)
		v2 := parseParam(p, p.Mem[p.Pointer+2], m2)
		v3 := parseAddr(p, p.Mem[p.Pointer+3], m3)

		switch opcode {
		case 1:
			p.Mem[v3] = v1 + v2
			p.Pointer += 4
		case 2:
			p.Mem[v3] = v1 * v2
			p.Pointer += 4
		case 3:
			if !p.Hang {
				v1 = parseAddr(p, p.Mem[p.Pointer+1], m1)
				p.Mem[v1] = in
				p.Pointer += 2
				p.Hang = true
			} else {
				p.Halts = true
			}
		case 4:
			p.Msg = append(p.Msg, v1)
			p.Pointer += 2
		case 5:
			if v1 != 0 {
				p.Pointer = v2
			} else {
				p.Pointer += 3
			}
		case 6:
			if v1 == 0 {
				p.Pointer = v2
			} else {
				p.Pointer += 3
			}
		case 7:
			if v1 < v2 {
				p.Mem[v3] = 1
			} else {
				p.Mem[v3] = 0
			}
			p.Pointer += 4
		case 8:
			if v1 == v2 {
				p.Mem[v3] = 1
			} else {
				p.Mem[v3] = 0
			}
			p.Pointer += 4
		case 9:
			p.Rbase += v1
			p.Pointer += 2
		case 99:
			p.Halts = true
		}
	}

	return p
}
