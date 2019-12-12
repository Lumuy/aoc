package main

import "testing"

func Test(t *testing.T) {
	n1, s1 := process(input_test1, 10)
	n2, s2 := process(input_test2, 100)

	if n1 != 179 {
		t.Errorf("Part one failed, should be 179, but %d", n1)
	}
	if n2 != 1940 {
		t.Errorf("Part one failed, should be 1940, but %d", n2)
	}
	if s1 != 2772 {
		t.Errorf("Part two failed, should be 2772, but %d", s1)
	}
	if s2 != 4686774924 {
		t.Errorf("Part two failed, should be 4686774924, but %d", s2)
	}
}

const input_test1 string = `<x=-1, y=0, z=2>
<x=2, y=-10, z=-7>
<x=4, y=-8, z=8>
<x=3, y=5, z=-1>`

const input_test2 string = `<x=-8, y=-10, z=0>
<x=5, y=5, z=10>
<x=2, y=-7, z=3>
<x=9, y=-8, z=-3>`
