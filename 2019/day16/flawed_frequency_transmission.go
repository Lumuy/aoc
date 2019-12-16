package main

import (
	"fmt"
	"strconv"
	"strings"
)

func getMessage(str string) (res string) {
	for p := 0; p < 100; p++ {
		newStr := ""
		for i, _ := range str {
			sum := 0
			for j, c := range str {
				sum += int(c-'0') * []int{0, 1, 0, -1}[(j+1)/(i+1)%4]
			}
			if sum < 0 {
				sum = -sum
			}
			newStr += string('0' + sum%10)
		}
		str = newStr
	}

	return str[:8]
}

func main() {
	{
		// Part 1
		s1 := getMessage(input)
		fmt.Println(s1)
	}

	{
		s2 := strings.Repeat(input, 10000)
		offset, _ := strconv.Atoi(s2[:7])
		output := []int{}

		for _, c := range s2[offset:] {
			output = append(output, int(c-'0'))
		}

		for p := 0; p < 100; p++ {
			sum := 0
			for i := len(output) - 1; i >= 0; i-- {
				sum += output[i]
				output[i] = sum % 10
			}
		}

		res := ""
		for _, c := range output[:8] {
			res += strconv.Itoa(c)
		}
		fmt.Println(res)
	}
}

const input string = `59765216634952147735419588186168416807782379738264316903583191841332176615408501571822799985693486107923593120590306960233536388988005024546603711148197317530759761108192873368036493650979511670847453153419517502952341650871529652340965572616173116797325184487863348469473923502602634441664981644497228824291038379070674902022830063886132391030654984448597653164862228739130676400263409084489497532639289817792086185750575438406913771907404006452592544814929272796192646846314361074786728172308710864379023028807580948201199540396460310280533771566824603456581043215999473424395046570134221182852363891114374810263887875638355730605895695123598637121`
