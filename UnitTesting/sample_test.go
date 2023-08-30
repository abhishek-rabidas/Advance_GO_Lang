package main

import (
	"fmt"
	"testing"
)

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

type TestCases struct {
	a, b int
	ans  int
}

func TestTableDriven(t *testing.T) {

	testCases := []TestCases{{
		a:   5,
		b:   10,
		ans: 10,
	},
		{
			a:   7,
			b:   4,
			ans: 7,
		},
		{
			a:   9,
			b:   20,
			ans: 20,
		}}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("max(%d, %d)", tc.a, tc.b), func(t *testing.T) {
			ans := getMaximum(tc.a, tc.b)
			if ans != tc.ans {
				t.Error("Wrong")
			}
		})
	}
}
