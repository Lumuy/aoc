package main

import (
	"fmt"
	"io/ioutil"
	"math/big"
	"regexp"
	"strconv"
	"strings"
)

const (
	Size  = int64(119315717514047)
	Times = int64(101741582076661)
)

const (
	KindNew       = 1
	KindCut       = 2
	KindIncrement = 3
)

type Instruction struct {
	Kind  int
	Value int64
}

func dealIntoNewStack(cards []int) (r []int) {
	for i := len(cards) - 1; i >= 0; i-- {
		r = append(r, cards[i])
	}
	return r
}

func cutCards(cards []int, n int) (r []int) {
	if n > 0 {
		r = append(cards[n:], cards[0:n]...)
	} else if n < 0 {
		s := len(cards) + n
		r = append(cards[s:], cards[:s]...)
	}
	return r
}

func dealWithIncrement(cards []int, n int) (r []int) {
	size := len(cards)
	// Grow result slice size with enough space
	for i := 0; i < size; i++ {
		r = append(r, 0)
	}

	ri := 0
	for i := 0; i < size; i++ {
		r[ri] = cards[i]
		ri = (ri + n) % size
	}

	return r
}

func process(inputs string, size int) []int {
	cards := genCards(size)

	for _, line := range strings.Split(inputs, "\n") {
		if strings.Contains(line, "new") {
			cards = dealIntoNewStack(cards)
		} else if strings.Contains(line, "cut") {
			n := parseNumber(line)
			cards = cutCards(cards, n)
		} else if strings.Contains(line, "increment") {
			n := parseNumber(line)
			cards = dealWithIncrement(cards, n)
		}
	}

	return cards
}

/*
	---------------------------Math explanation-----------------------------
	functions:

	to deal into new stack   => fa(x)
	to cut N card            => fb(x)
	to deal with increment N => fc(x)

	x: index
	size: number of cards

	fa(x) = size - 1 - x
	fb(x) = (x + N + size) % size
	fc(x) = (N % size) * x % size

	these functions are linear, all can repsent by:
		f(x) = A * x + B

	for a, b, c, b, a, c type steps:
		y = fa(fb(fc(fb(fa(fc(x)))))) = A * x + B
	here a, b is determined by the a, b, c, b, a, c changeless inputs

	for three cycles:
		y = A*(A*(A*x + B) + B) + B
		  = A^3*x + A^2*B + A*B + B
			= A^3*x + (A^2 + A + 1)*b

	in general:
		f^n(x) = A^n*x + A^(n-1)*B + A^(n-2)*B + ... + B
		       = A^n*x + (A^(n-1) + A^(n-2) + ... + 1)*B
					 = A^n*x + (A^n - 1) / (A - 1) * B

	here n = 101741582076661

	reverse function:
		pow(A, n, size)*X + (pow(A, n, size)-1) / modinv(A-1, size) * B) % size
	------------------------------------------------------------------------
*/

/*
	Two consecutive "deal into stack" will cancel each other.

					LEFT            |         RIGHT
	-------------------------------------------------
	| deal into new stack   | cut size-x            |
	| cut x                 | deal into new stack   |
	-------------------------------------------------
	| deal into new stack   | deal with increment x |
	| deal with increment x | cut size+1-x          |
	|												| deal into new stack   |
	-------------------------------------------------

	-------------------------------------------------
	| cut x                 | cut (x+y)%size        |
	| cut y                 |                       |
	-------------------------------------------------
	| cut x                 | deal with increment y |
	| deal with increment y | cut (x*y)%szie        |
	-------------------------------------------------

	-----------------------------------------------------------
	| deal with increment x | deal with increment (x*y)%size  |
	| deal with increment y |															    |
	----------------------------------------------------------

	The table's left has same effect with right.
*/
func compactInputs(filename string) []Instruction {
	var ins1, ins2, ins3, ins4 []Instruction

	text := readFile(filename)
	for _, line := range strings.Split(text, "\n") {
		ins1 = append(ins1, parseLine(line))
	}

	fmt.Println("-----------------")
	fmt.Println(len(ins1))
	fmt.Println(ins1)
	for i := 0; i < len(ins1)-1; {
		in1, in2 := ins1[i], ins1[i+1]
		if in1.Kind == KindNew && in2.Kind == KindCut {
			add := []Instruction{{KindCut, Size - in2.Value}, in1}
			ins2 = append(ins2, add...)
			i += 2
		} else if in1.Kind == KindNew && in2.Kind == KindIncrement {
			add := []Instruction{in2, {KindCut, Size + 1 - in2.Value}, in1}
			ins2 = append(ins2, add...)
			i += 2
		} else {
			ins2 = append(ins2, in1)
			i++
		}
	}
	fmt.Println("-----------------")
	fmt.Println(len(ins2))
	fmt.Println(ins2)
	for i := 0; i < len(ins2)-1; {
		in1, in2 := ins2[i], ins2[i+1]
		if in1.Kind == KindNew && in2.Kind == KindNew {
			fmt.Println("consecutive reverse: =>", in1, in2)
			i += 2
		} else if in1.Kind == KindCut && in2.Kind == KindCut {
			add := Instruction{KindCut, (in1.Value + in2.Value) % Size}
			ins3 = append(ins3, add)
			i += 2
		} else {
			ins3 = append(ins3, in1)
			i++
		}
	}
	fmt.Println("-----------------")
	fmt.Println(len(ins3))
	fmt.Println(ins3)

	for i := 0; i < len(ins3)-1; {
		in1, in2 := ins3[i], ins3[i+1]
		if in1.Kind == KindNew && in2.Kind == KindNew {
			fmt.Println("consecutive reverse: =>", in1, in2)
			i += 2
		} else if in1.Kind == KindIncrement && in2.Kind == KindIncrement {
			add := Instruction{KindIncrement, modMul(in1.Value, in2.Value)}
			ins4 = append(ins4, add)
			i += 2
		} else {
			ins4 = append(ins4, in1)
		}
	}
	fmt.Println("-----------------")
	fmt.Println(len(ins4))
	fmt.Println(ins4)

	return ins4
}

func modMul(l, r int64) int64 {
	a, b, s := big.NewInt(l), big.NewInt(r), big.NewInt(Size)
	res := big.NewInt(0)
	res = res.Mul(a, b)
	res = res.Mod(res, s)

	return res.Int64()
}

func parseLine(line string) Instruction {
	var kind int
	var value int64

	if strings.Contains(line, "new") {
		kind = KindNew
	} else if strings.Contains(line, "cut") {
		kind = KindCut
	} else if strings.Contains(line, "increment") {
		kind = KindIncrement
	} else {
		panic(line)
	}

	if kind == KindCut || kind == KindIncrement {
		re := regexp.MustCompile(`[-]?\d+`)
		rs := re.FindString(line)
		value, _ = strconv.ParseInt(rs, 10, 64)
	}

	return Instruction{kind, value}
}

func main() {
	inputs := readFile("input")
	cards := process(inputs, 10007)

	// Part 1
	for i, v := range cards {
		if v == 2019 {
			fmt.Println(i)
			break
		}
	}

	// Part 2
	compactInputs("input")
}

func parseNumber(s string) int {
	re := regexp.MustCompile(`[-]?\d+`)
	rs := re.FindString(s)
	n, err := strconv.Atoi(rs)
	check(err)

	return n
}

func readFile(filename string) string {
	bytes, err := ioutil.ReadFile(filename)
	check(err)
	return strings.TrimSpace(string(bytes))
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func genCards(size int) (r []int) {
	for i := 0; i < size; i++ {
		r = append(r, i)
	}
	return r
}
