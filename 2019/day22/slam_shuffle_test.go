package main

import (
	"strconv"
	"strings"
	"testing"
)

func convert(data []int) string {
	t := make([]string, 0)
	for _, n := range data {
		s := strconv.Itoa(n)
		t = append(t, s)
	}
	return strings.Join(t, " ")
}

func Test(t *testing.T) {
	r1 := process(input1, 10)
	r2 := process(input2, 10)
	r3 := process(input3, 10)
	r4 := process(input4, 10)

	s1 := convert(r1)
	s2 := convert(r2)
	s3 := convert(r3)
	s4 := convert(r4)

	if s1 != "0 3 6 9 2 5 8 1 4 7" {
		t.Errorf("Failed test 1, should be , but be %s", s1)
	}
	if s2 != "3 0 7 4 1 8 5 2 9 6" {
		t.Errorf("Failed test 2, should be , but be %s", s2)
	}
	if s3 != "6 3 0 7 4 1 8 5 2 9" {
		t.Errorf("Failed test 3, should be , but be %s", s3)
	}
	if s4 != "9 2 5 8 1 4 7 0 3 6" {
		t.Errorf("Failed test 4, should be , but be %s", s4)
	}
}

const input1 string = `
deal with increment 7
deal into new stack
deal into new stack
`

const input2 string = `
cut 6
deal with increment 7
deal into new stack
`

const input3 string = `
deal with increment 7
deal with increment 9
cut -2
`

const input4 string = `
deal into new stack
cut -2
deal with increment 7
cut 8
cut -4
deal with increment 7
cut 3
deal with increment 9
deal with increment 3
cut -1
`
