package main

import (
	"encoding/base64"
	"fmt"
)

func main() {

	fmt.Println("Enter the string to encode into base64: ")
	var data string
	_, err := fmt.Scanln(&data)
	if err != nil {
		panic(err)
	}

	encodedString := base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println("Encoded string: ", encodedString)

	// URL compatible encoding
	encodedString = base64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println("Encoded string: ", encodedString)
}
