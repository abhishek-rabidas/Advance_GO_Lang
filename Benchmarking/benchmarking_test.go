package main

import (
	"fmt"
	"github.com/labstack/gommon/random"
	"testing"
)

func reverseString() {
	str := random.Alphanumeric
	runes := []rune(str)
	length := len(runes)
	for i := 0; i < length/2; i++ {
		runes[i], runes[length-i-1] = runes[length-i-1], runes[i]
	}

	fmt.Println("Actual string: ", str)
	fmt.Println("Reversed string: ", string(runes))
}

/*
Benchmark tests typically go in _test.go files and are named beginning with Benchmark. The testing runner executes each
benchmark function several times, increasing b.N on each run until it collects a precise measurement.
*/

func BenchmarkRevString(t *testing.B) {

	for i := 0; i < t.N; i++ {
		reverseString()
	}
}
