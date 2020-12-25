package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func parseInput(filename string) (r []int) {
	dat, _ := ioutil.ReadFile(filename)
	for _, line := range strings.Split(strings.TrimSpace(string(dat)), "\n") {
		n, _ := strconv.Atoi(line)
		r = append(r, n)
	}
	return r
}

func getLoopSize(pk int) (size int) {
	val, sn := 1, 7
	for {
		size++
		val = (val * sn) % 20201227
		if val == pk {
			break
		}
	}
	return size
}

func getEncryptionKey(sn, size int) int {
	val := 1
	for i := 0; i < size; i++ {
		val = (val * sn) % 20201227
	}
	return val
}

func run(pks []int) int {
	return getEncryptionKey(pks[0], getLoopSize(pks[1]))
}

func main() {
	pks := parseInput("input")
	fmt.Println(run(pks))
}
