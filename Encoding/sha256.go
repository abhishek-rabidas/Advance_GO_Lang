package main

import "crypto/sha256"

func main() {
	str := "this is a string"

	encoder := sha256.New()

	encoder.Write([]byte(str))
}
