package intcode_program

func MaxInt(x1, x2 int) int {
	if x1 < x2 {
		return x2
	}
	return x1
}

func MinInt(x1, x2 int) int {
	if x1 > x2 {
		return x2
	}
	return x1
}
