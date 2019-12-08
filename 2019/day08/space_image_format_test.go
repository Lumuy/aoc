package main

import "testing"

func Test(t *testing.T) {
	n1 := getFewestLayerNumber(input_test1, 3, 2)
	if n1 != 1 {
		t.Errorf("Test 1 failed, should be 1, but %d", n1)
	}
}

const input_test1 string = `123456789012`
