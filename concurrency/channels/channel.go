package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan string)

	go write("Hello world", channel)

	fmt.Println("After")

	for {
		message, open := <-channel
		fmt.Println(message)

		if !open {
			break
		}
	}

	fmt.Println("End")
}

func write(text string, ch chan string) {
	time.Sleep(time.Second * 5)
	for i := 0; i < 5; i++ {
		ch <- text
		time.Sleep(time.Second)
	}
	close(ch)
}
