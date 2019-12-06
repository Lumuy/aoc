package main

import "testing"

func TestGetDiagnosticCode(t *testing.T) {
	if getDiagnosticCode(input_test, 1) != 999 {
		t.Error("Failed on id 1")
	}
	if getDiagnosticCode(input_test, 2) != 999 {
		t.Error("Failed on id 2")
	}
	if getDiagnosticCode(input_test, 3) != 999 {
		t.Error("Failed on id 3")
	}
	if getDiagnosticCode(input_test, 4) != 999 {
		t.Error("Failed on id 4")
	}
	if getDiagnosticCode(input_test, 5) != 999 {
		t.Error("Failed on id 5")
	}
	if getDiagnosticCode(input_test, 6) != 999 {
		t.Error("Failed on id 6")
	}
	if getDiagnosticCode(input_test, 7) != 999 {
		t.Error("Failed on id 7")
	}
	if getDiagnosticCode(input_test, 8) != 1000 {
		t.Error("Failed on id 8")
	}
	if getDiagnosticCode(input_test, 9) != 1001 {
		t.Error("Failed on id 9")
	}
	if getDiagnosticCode(input_test, 10) != 1001 {
		t.Error("Failed on id 10")
	}
	if getDiagnosticCode(input_test, 11) != 1001 {
		t.Error("Failed on id 11")
	}
	if getDiagnosticCode(input_test, 12) != 1001 {
		t.Error("Failed on id 12")
	}
}

const input_test string = `3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104,999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99`
