package main

import "fmt"

func main() {

	fmt.Println("Enter the string to encode into base64: ")
	var data string
	_, err := fmt.Scanln(&data)
	if err != nil {
		return
	}

}
