package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan string)
	go requester(channel)
	go receiver(channel)

	fmt.Scanln()
}

func requester(channel chan string) {
	for {
		channel <- fmt.Sprintf("MSG [%s]", time.Now().String())
	}
}

func receiver(channel chan string) {
	for {
		fmt.Println(<-channel)
	}
}
