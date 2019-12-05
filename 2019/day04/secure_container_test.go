package main

import "testing"

func TestIsNotPartOfLargerGroup(t *testing.T) {
	if !isNotPartOfLargerGroup(112233) {
		t.Error("112233 failed")
	}
	if isNotPartOfLargerGroup(123444) {
		t.Error("123444 failed")
	}
	if !isNotPartOfLargerGroup(111122) {
		t.Error("111122 failed")
	}
}
