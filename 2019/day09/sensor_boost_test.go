package main

import "testing"

func Test(t *testing.T) {
	code1 := process(getInput(input_test1), 1)
	code2 := process(getInput(input_test2), 1)
	code3 := process(getInput(input_test3), 1)
	if code1 != "1125899906842624" {
		t.Errorf("Test failed, should be 1125899906842624, but %s", code1)
	}
	if code2 != "1091204-1100110011001008100161011006101099" {
		t.Errorf("Test failed, should be %s, but %s", "1091204-1100110011001008100161011006101099", code2)
	}
	if code3 != "1219070632396864" {
		t.Errorf("Test failed, should be 1219070632396864, but %s", code3)
	}
}

const input_test1 string = `104,1125899906842624,99`
const input_test2 string = `109,1,204,-1,1001,100,1,100,1008,100,16,101,1006,101,0,99`
const input_test3 string = `1102,34915192,34915192,7,4,7,99,0`
