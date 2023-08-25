package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(2 * time.Second)

	go func() {
		for {
			<-ticker.C
			fmt.Println("Ticker received")
		}
	}()
	time.Sleep(1 * time.Hour) // halting the program from exiting
}
