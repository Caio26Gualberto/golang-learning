package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	channel := multiplexator(write("Heelo World"), write("Golang"))

	for i := 0; i < 10; i++ {
		fmt.Println(<-channel)
	}
}

func multiplexator(channel1, channel2 <-chan string) chan string {
	outputChannel := make(chan string)

	go func() {
		for {
			select {
			case message := <-channel1:
				outputChannel <- message
			case message := <-channel2:
				outputChannel <- message
			}
		}
	}()

	return outputChannel
}

func write(text string) <-chan string {
	channel := make(chan string)

	go func() {
		for {
			channel <- fmt.Sprintf("Valor recebido: %s", text)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(2000)))
		}
	}()

	return channel
}
