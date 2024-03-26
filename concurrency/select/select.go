package main

import (
	"fmt"
	"time"
)

func main() {
	channel, channel2 := make(chan string), make(chan string)

	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			channel <- "Canal 1"

		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 2)
			channel2 <- "Canal 2"

		}
	}()

	for {
		select {
		case messsageChannel1 := <-channel:
			fmt.Println(messsageChannel1)
		case messsageChannel2 := <-channel2:
			fmt.Println(messsageChannel2)
		}
	}

}
