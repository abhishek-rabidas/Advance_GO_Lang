package main

import (
	"fmt"
	"time"
)

func main() {
	messageChannel := make(chan string)

	//async function to store messages into channel
	go func() {
		for i := 0; i < 5; i++ {
			messageChannel <- "time: " + time.Now().String()
		}
		close(messageChannel) //closing the channel
	}()

	//getting messages from channel and printing
	for {
		msg := <-messageChannel

		if msg == "" {
			break
		}

		fmt.Println(msg)
	}

	fmt.Println("END")

}
