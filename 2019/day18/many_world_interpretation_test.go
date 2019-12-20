package main

import "testing"

func Test(t *testing.T) {
	s1 := shortestPathSteps(input1)
	s2 := shortestPathSteps(input2)
	s3 := shortestPathSteps(input3)

	if s1 != 132 {
		t.Errorf("Part 1 failed, should be 132, but be %d", s1)
	}
	if s2 != 136 {
		t.Errorf("Part 1 failed, should be 136, but be %d", s2)
	}
	if s3 != 81 {
		t.Errorf("Part 1 failed, should be 81, but be %d", s3)
	}
}

const input1 string = `
########################
#...............b.C.D.f#
#.######################
#.....@.a.B.c.d.A.e.F.g#
########################
`

const input2 string = `
#################
#i.G..c...e..H.p#
########.########
#j.A..b...f..D.o#
########@########
#k.E..a...g..B.n#
########.########
#l.F..d...h..C.m#
#################
`

const input3 string = `
########################
#@..............ac.GI.b#
###d#e#f################
###A#B#C################
###g#h#i################
########################
`
