package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type Passport struct {
	pid, ecl, hcl, hgt string
	eyr, byr, iyr, cid int
}

func (pt *Passport) isValid() bool {
	valid := pt.eyr*pt.byr*pt.iyr != 0
	if valid && pt.ecl != "" && pt.hcl != "" && pt.hgt != "" && pt.pid != "" {
		return true
	}
	return false
}

func (pt *Passport) isValidHgt() bool {
	ar := regexp.MustCompile(`(\d+)(.*)`).FindStringSubmatch(pt.hgt)

	if len(ar) != 3 {
		return false
	}
	hgt, _ := strconv.Atoi(ar[1])

	switch ar[2] {
	case "cm":
		return hgt >= 150 && hgt <= 193
	case "in":
		return hgt >= 59 && hgt <= 76
	}
	return false
}

func (pt *Passport) isValidHcl() bool {
	return regexp.MustCompile(`^#[\da-f]{6}$`).MatchString(pt.hcl)
}

func (pt *Passport) isValidEcl() bool {
	for _, color := range []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"} {
		if color == pt.ecl {
			return true
		}
	}
	return false
}

func (pt *Passport) isValidPid() bool {
	re := regexp.MustCompile(`^[0-9]{9}$`)
	return re.MatchString(pt.pid)
}

func (pt *Passport) isValidByr() bool {
	return pt.byr >= 1920 && pt.byr <= 2002
}

func (pt *Passport) isValidIyr() bool {
	return pt.iyr >= 2010 && pt.iyr <= 2020
}

func (pt *Passport) isValidEyr() bool {
	return pt.eyr >= 2020 && pt.eyr <= 2030
}

func (pt *Passport) isStrictValid() bool {
	valid := pt.isValidPid() && pt.isValidEcl() && pt.isValidHcl() && pt.isValidHgt()
	return valid && pt.isValidByr() && pt.isValidIyr() && pt.isValidEyr()
}

func getPassports(filename string) (r []Passport) {
	dat, _ := ioutil.ReadFile(filename)
	for _, line := range strings.Split(string(dat), "\n\n") {
		if line != "" {
			var pt Passport
			line = strings.ReplaceAll(line, "\n", " ")
			for _, ele := range strings.Split(line, " ") {
				if ele != "" {
					kv := strings.Split(ele, ":")
					switch kv[0] {
					case "ecl":
						pt.ecl = kv[1]
					case "hcl":
						pt.hcl = kv[1]
					case "hgt":
						pt.hgt = kv[1]
					case "pid":
						pt.pid = kv[1]
					case "eyr":
						n, _ := strconv.Atoi(kv[1])
						pt.eyr = n
					case "byr":
						n, _ := strconv.Atoi(kv[1])
						pt.byr = n
					case "iyr":
						n, _ := strconv.Atoi(kv[1])
						pt.iyr = n
					case "cid":
						n, _ := strconv.Atoi(kv[1])
						pt.cid = n
					}
				}
			}
			r = append(r, pt)
		}
	}
	return r
}

func countValidPassports(passports []Passport) (r int) {
	for _, pt := range passports {
		if pt.isValid() {
			r++
		}
	}
	return r
}

func countValidPassportsTwo(passports []Passport) (r int) {
	for _, pt := range passports {
		if pt.isValid() && pt.isStrictValid() {
			r++
		}
	}
	return r
}

func main() {
	passports := getPassports("input")
	fmt.Println(countValidPassports(passports))
	fmt.Println(countValidPassportsTwo(passports))
}
