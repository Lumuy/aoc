package main

import (
	"fmt"
)

type Node struct {
	val  int
	prev *Node
	next *Node
}

func (p *Node) pickUp(n int) (rp, ra *Node, rm map[int]bool) {
	rm = map[int]bool{}
	tp := p
	for i := 0; i < n; i++ {
		ra = p.next
		if i == 0 {
			rp = ra
		}
		rm[ra.val] = true
		p = p.next
	}
	np := ra.next
	tp.next = np
	np.prev = tp

	return rp, ra, rm
}

func (p *Node) insert(rp, ra *Node) {
	pn := p.next
	p.next = rp
	rp.prev = p
	pn.prev = ra
	ra.next = pn
}

func newCircleList(labels string, ncups int) (p *Node, m map[int]*Node) {
	var sp, cp *Node
	m = map[int]*Node{}

	for i, r := range labels {
		np := &Node{val: int(r - '0')}
		m[int(r-'0')] = np
		if i == 0 {
			sp = np
			cp = np
		} else {
			np.prev = cp
			cp.next = np
		}
		cp = np
	}
	if ncups > len(labels) {
		for i := 10; i <= ncups; i++ {
			np := &Node{val: i}
			m[i] = np
			cp.next = np
			np.prev = cp
			cp = np
		}
	}
	cp.next = sp
	sp.prev = cp

	return sp, m
}

func play(labels string, ncups int, moves int) (rs string) {
	lp, lm := newCircleList(labels, ncups)

	for i := 0; i < moves; i++ {
		des := (ncups+lp.val-2)%ncups + 1
		pb, pa, banned := lp.pickUp(3)
		for banned[des] {
			des = (ncups+des-2)%ncups + 1
		}
		lm[des].insert(pb, pa)
		lp = lp.next
	}

	p := lm[1]
	if ncups > len(labels) {
		rs = fmt.Sprintf("%d", p.next.next.val*p.next.val)
	} else {
		for {
			p = p.next
			if p.val == 1 {
				break
			}
			rs += fmt.Sprintf("%d", p.val)
		}
	}

	return rs
}

func main() {
	labels := "583976241"
	str1 := play(labels, len(labels), 100)
	str2 := play(labels, 1000000, 10000000)
	fmt.Println(str1, str2)
}
