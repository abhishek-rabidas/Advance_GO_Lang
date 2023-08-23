package main

import (
	"fmt"
	"time"
)

func main() {
	messageChannel := make(chan string)

	go func() {
		for {
			messageChannel <- "time: " + time.Now().String()
		}
	}()

	for {
		msg := <-messageChannel

		fmt.Println(msg)
	}

}
