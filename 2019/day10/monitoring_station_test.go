package main

import "testing"

func Test(t *testing.T) {
	// Part 1 test
	_, n1 := getMaxDetectedAsteroid(input_test1)
	_, n2 := getMaxDetectedAsteroid(input_test2)
	_, n3 := getMaxDetectedAsteroid(input_test3)
	_, n4 := getMaxDetectedAsteroid(input_test4)
	l, n5 := getMaxDetectedAsteroid(input_test5)

	if n1 != 8 {
		t.Errorf("Test 1 failed, should be 8, but %d", n1)
	}
	if n2 != 33 {
		t.Errorf("Test 1 failed, should be 33, but %d", n2)
	}
	if n3 != 35 {
		t.Errorf("Test 1 failed, should be 35, but %d", n3)
	}
	if n4 != 41 {
		t.Errorf("Test 1 failed, should be 41, but %d", n4)
	}
	if n5 != 210 {
		t.Errorf("Test 1 failed, should be 210, but %d", n5)
	}
	if (l != Point{11, 13}) {
		t.Errorf("Test failed, laser should be 11, 13, but %f, %f", l.x, l.y)
	}

	// Part 2 test
	p1 := getVaporizedPoint(input_test5, 3)
	p2 := getVaporizedPoint(input_test5, 50)
	p3 := getVaporizedPoint(input_test5, 100)
	p4 := getVaporizedPoint(input_test5, 201)
	p5 := getVaporizedPoint(input_test5, 299)

	if (p1 != Point{12, 2}) {
		t.Errorf("Test 2 failed, should be x: 12.0 == %f, y: 2.0 == %f", p1.x, p1.y)
	}
	if (p2 != Point{16, 9}) {
		t.Errorf("Test 2 failed, should be x: 16.0 == %f, y: 9.0 == %f", p2.x, p2.y)
	}
	if (p3 != Point{10, 16}) {
		t.Errorf("Test 2 failed, should be x: 10.0 == %f, y: 16.0 == %f", p3.x, p3.y)
	}
	if (p4 != Point{10, 9}) {
		t.Errorf("Test 2 failed, should be x: 10.0 == %f, y: 9.0 == %f", p4.x, p4.y)
	}
	if (p5 != Point{11, 1}) {
		t.Errorf("Test 2 failed, should be x: 11.0 == %f, y: 1.0 == %f", p5.x, p5.y)
	}
}

const input_test1 string = `.#..#
.....
#####
....#
...##`

const input_test2 string = `......#.#.
#..#.#....
..#######.
.#.#.###..
.#..#.....
..#....#.#
#..#....#.
.##.#..###
##...#..#.
.#....####`

const input_test3 string = `#.#...#.#.
.###....#.
.#....#...
##.#.#.#.#
....#.#.#.
.##..###.#
..#...##..
..##....##
......#...
.####.###.`

const input_test4 string = `.#..#..###
####.###.#
....###.#.
..###.##.#
##.##.#.#.
....###..#
..#.#..#.#
#..#.#.###
.##...##.#
.....#.#..`

const input_test5 string = `.#..##.###...#######
##.############..##.
.#.######.########.#
.###.#######.####.#.
#####.##.#.##.###.##
..#####..#.#########
####################
#.####....###.#.#.##
##.#################
#####.##.###..####..
..######..##.#######
####.##.####...##..#
.#####..#.######.###
##...#.##########...
#.##########.#######
.####.#.###.###.#.##
....##.##.###..#####
.#.#.###########.###
#.#.#.#####.####.###
###.##.####.##.#..##`
