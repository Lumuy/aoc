package main

import "testing"

func Test(t *testing.T) {
	n1 := getMessage(input_test1)
	n2 := getMessage(input_test2)
	n3 := getMessage(input_test3)

	if n1 != `24176176` {
		t.Errorf("After 100 pahses, should be 24176176, but be %s", n1)
	}
	if n2 != `73745418` {
		t.Errorf("After 100 pahses, should be 73745418, but be %s", n2)
	}
	if n3 != `52432133` {
		t.Errorf("After 100 pahses, should be 52432133, but be %s", n3)
	}
}

const input_test1 string = `80871224585914546619083218645595`
const input_test2 string = `19617804207202209144916044189917`
const input_test3 string = `69317163492948606335995924319873`
