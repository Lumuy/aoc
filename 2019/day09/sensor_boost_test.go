package main

import "testing"

func Test(t *testing.T) {
	p1 := process(input_test1, 1)
	p2 := process(input_test2, 1)
	p3 := process(input_test3, 1)
	if p1.Msg[0] != 1125899906842624 {
		t.Errorf("Test failed, should be 1125899906842624, but %d", p1.Msg[0])
	}
	if p2.Msg[0] != 109 {
		t.Errorf("Test failed, should be %d, but %d", 109, p2.Msg[0])
	}
	if p3.Msg[0] != 1219070632396864 {
		t.Errorf("Test failed, should be 1219070632396864, but %d", p3.Msg[0])
	}
}

const input_test1 string = `104,1125899906842624,99`
const input_test2 string = `109,1,204,-1,1001,100,1,100,1008,100,16,101,1006,101,0,99`
const input_test3 string = `1102,34915192,34915192,7,4,7,99,0`
