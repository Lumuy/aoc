package main

import "testing"

func Test(t *testing.T) {
	nts := parseNotes("test_input")
	_, ert := getTicketScanningErrorRate(nts)

	if ert != 71 {
		t.Errorf("Yout ticket scanning error rate should be 71, but be %d", ert)
	}
}
