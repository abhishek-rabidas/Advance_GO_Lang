package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	encodedHash := "YWJoaXNoZWs="

	decodedStr, _ := base64.StdEncoding.DecodeString(encodedHash)
	fmt.Println("Decoded string: ", string(decodedStr))
}
