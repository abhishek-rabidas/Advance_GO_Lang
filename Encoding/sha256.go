package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	str := "this is a string"

	encoder := sha256.New()

	encoder.Write([]byte(str))
	bs := encoder.Sum(nil)

	fmt.Println("Encoded String: ", bs)
}
