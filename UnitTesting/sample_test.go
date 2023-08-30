package main

import "testing"

func getMaximum(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func basicTest(t *testing.T) {
	ans := getMaximum(2, 5)
	if ans != 5 {
		t.Error("Wrong")
	}
}
