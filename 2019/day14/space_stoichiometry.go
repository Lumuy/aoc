package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Unit struct {
	name  string
	count int
}

func getInput(in string) map[Unit][]Unit {
	reactions := make(map[Unit][]Unit)

	for _, line := range strings.Split(in, "\n") {
		key := Unit{}
		val := []Unit{}

		re := regexp.MustCompile(`\d+\s\w+`)
		arr := re.FindAllString(line, -1)
		for i, str := range arr {
			inputs := strings.Split(str, " ")
			count, _ := strconv.Atoi(inputs[0])
			unit := Unit{name: inputs[1], count: count}

			if i == len(arr)-1 {
				key = unit
			} else {
				val = append(val, unit)
			}
		}

		reactions[key] = val
	}

	return reactions
}

func getReactionsOutput(data map[Unit][]Unit, key Unit, remain map[string]int) (res []Unit, newRemain map[string]int, status bool) {
	base := 0

	for k, v := range data {
		if k.name == key.name {
			more := key.count % k.count
			if more != 0 {
				if remain[k.name] >= more {
					remain[k.name] -= more
					base = key.count / k.count
				} else {
					remain[k.name] = k.count - more + remain[k.name]
					base = key.count/k.count + 1
				}
			} else {
				base = key.count / k.count
			}

			for _, u := range v {
				unit := Unit{u.name, u.count * base}
				res = append(res, unit)
			}
			status = true
			break
		}
	}

	newRemain = remain

	return res, newRemain, status
}

func fuelRequiredOre(in string, number int) int {
	reactions := getInput(in)
	res := []Unit{{"FUEL", number}}
	remain := make(map[string]int)
	goon := true

	for goon {
		goon = false
		tmp := make(map[string]int)
		for _, input := range res {
			units, newRemain, ok := getReactionsOutput(reactions, input, remain)
			remain = newRemain
			goon = goon || ok

			if ok {
				for _, unit := range units {
					tmp[unit.name] += unit.count
				}
			} else {
				tmp[input.name] += input.count
			}
		}

		res = []Unit{}
		for name, count := range tmp {
			res = append(res, Unit{name, count})
		}
	}

	return res[0].count
}

func maximumProducedFuel(in string) int {
	base := fuelRequiredOre(in, 1)
	limit := 1000000000000

	for i := limit / base; ; {
		count := fuelRequiredOre(in, i)
		if count < limit {
			if base == 0 {
				return i
			}
			i += base
		} else if count > limit {
			i -= base
			base = base / 2
		} else {
			return i
		}
	}
}

func main() {
	// Part 1
	fmt.Println(fuelRequiredOre(input, 1))
	// Part 2
	fmt.Println(maximumProducedFuel(input))
}

const input string = `3 JQFM, 5 QMQB, 20 WQCT => 8 PHBMP
2 XJFQR => 1 WQCT
133 ORE => 3 KFKWH
1 QGVJV, 9 TNRGW, 9 NSWDH => 5 HJPD
4 QMQB, 2 QZMZ, 3 DQPX, 1 HJFV, 5 SLQN, 4 XHKG, 23 DBKL => 5 CVZLJ
6 GFDP, 1 MXQF => 7 TDPN
19 BWHKF, 2 KXHQW, 8 GHNG, 8 CSNS, 8 JVRQ, 1 PHBMP, 20 LZWR, 7 JKRZH => 5 PZRSQ
1 JQFM => 1 QGVJV
8 KFKWH => 7 ZJKB
3 VMDSG, 2 BMSNV => 9 XJFQR
7 ZKZHV => 6 DVRS
1 WKFTZ, 5 SVTK, 1 QDJD => 7 JQFM
19 FRTK => 2 QMTMN
23 DVRS, 3 XNGTQ => 9 MCWF
188 ORE => 3 HDRMK
3 MCWF => 5 LHXN
12 KFKWH, 2 DWBL => 8 ZKZHV
2 GHNG => 8 SVTK
4 MLJN, 9 CSNS => 6 DQPX
2 NDNP, 1 LWGNJ, 6 DBKL, 3 RLKDF, 9 DQPX, 1 BWHKF => 3 JVGRC
4 TNRGW => 2 CFBP
2 KXHQW => 1 BWHKF
7 HJFV => 4 PDKZ
2 QZMZ => 5 BMSNV
1 SVTK, 1 LZWR, 1 WQCT => 3 SLQN
1 TDPN, 1 VMDSG => 7 NHVQD
1 SVTK => 2 LZWR
149 ORE => 9 DWBL
1 XMHN, 1 PDKZ => 5 LWGNJ
6 WCMV => 6 XNGTQ
7 MCWF, 2 VCMPS => 1 HJFV
11 BRTX, 37 CFBP, 2 HJPD, 72 HDRMK, 5 LWGNJ, 7 JVGRC, 3 CVZLJ, 8 PZRSQ, 3 LQBJP => 1 FUEL
9 QMTMN, 14 FRTK, 14 HJFV => 9 NDNP
1 KFKWH, 3 ZJKB => 9 MXQF
1 HJFV, 1 XJFQR => 9 TNRGW
1 DVRS => 2 BRTX
4 QZMZ, 3 BMSNV, 3 GFDP => 6 VMDSG
3 NHVQD => 6 WKFTZ
1 BWHKF => 6 DBKL
8 DWBL => 8 QZMZ
4 MLJN, 16 NSWDH, 4 XHKG => 8 JVRQ
2 DVRS, 32 XNGTQ, 9 MXQF => 7 GHNG
1 DWBL => 8 WCMV
8 SLQN, 1 CFBP => 9 MLJN
1 QDJD => 4 XMHN
3 BWHKF, 2 TNRGW => 9 XHKG
1 WGLN => 8 GFDP
1 MCWF, 1 XJFQR => 2 CSNS
3 XNGTQ => 1 QDJD
15 KXHQW, 3 WQCT, 2 QMTMN => 8 NSWDH
9 XCMJ, 1 QMTMN => 2 JKRZH
4 HDRMK => 4 WGLN
9 NSWDH, 12 LHXN, 16 NDNP => 1 QMQB
16 NHVQD, 3 DWBL, 1 WKFTZ => 4 FRTK
1 GFDP => 2 VCMPS
2 JQFM, 2 XMHN => 6 XCMJ
1 DVRS, 19 QZMZ, 1 DWBL => 5 KXHQW
1 QGVJV, 8 NDNP, 5 PDKZ => 1 RLKDF
29 HJFV, 2 WKFTZ, 4 GFDP => 2 LQBJP`
