package main

import (
	"fmt"
	"time"
)

func main() {
	go write("Olá mundo")
	write("Programming in go")
}

func write(text string) {
	for {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}
