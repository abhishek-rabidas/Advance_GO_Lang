package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func main() {
	str := "this is a string"

	encoder := sha256.New()

	encoder.Write([]byte(str))
	bs := encoder.Sum(nil)

	fmt.Println("Encoded String: ", bs)
	fmt.Printf("%x\n", bs)

	//compare hash
	hash := "bc7e8a24e2911a5827c9b33d618531ef094937f2b3803a591c625d0ede1fffc6"
	hexString := hex.EncodeToString(bs)

	if hexString == hash {
		fmt.Println("Hashes are matching")
	} else {
		fmt.Println("Hashes are not matching")
	}
}

// another way of SHA256 encoding
func directSHA256Encoding() {
	input := "Hello, world!"
	hash := sha256.Sum256([]byte(input))

	hashString := hex.EncodeToString(hash[:])
	fmt.Println("SHA-256 Hash:", hashString)
}
