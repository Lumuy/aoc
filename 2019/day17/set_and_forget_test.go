package main

import "testing"

func Test(t *testing.T) {
	movement := "R,8,R,8,R,4,R,4,R,8,L,6,L,2,R,4,R,4,R,8,R,8,R,8,L,6,L,2"
	res := compressInputs(movement)

	if inputs != res {
		t.Errorf("inputs should be %s \n, but be %s", inputs, res)
	}
}

const inputs string = `65,44,66,44,67,44,66,44,65,44,67,10,82,44,56,44,82,44,56,10,82,44,52,44,82,44,52,44,82,44,56,10,76,44,54,44,76,44,50,10`
