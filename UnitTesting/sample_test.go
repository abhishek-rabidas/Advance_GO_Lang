package main

import "testing"

// for a test file we need to postfix name the file _test, test function in Go includes this signature: func TestXxxx(t *testing.T) and then run the command go test in the package

func getMaximum(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// basic test method
func TestBasic(t *testing.T) {
	ans := getMaximum(2, 5)
	if ans != 5 {
		t.Error("Wrong")
	}
}

// table driven method
func TestTableDriven(t *testing.T) {

}
