package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func fewestSteps(inputs string) (num int) {
	inputs = strings.Trim(inputs, "\n")
	fmt.Println(inputs)

	return num
}

func main() {
	dat, _ := ioutil.ReadFile("input")
	inputs := string(dat)

	{
		// Part 1
		steps := fewestSteps(inputs)
		fmt.Println(steps)
	}
}
