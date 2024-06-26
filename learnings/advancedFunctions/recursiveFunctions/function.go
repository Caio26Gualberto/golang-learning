package main

import "fmt"

func main() {
	position := uint(12)
	for i := uint(0); i < position; i++ {
		fmt.Println(fibonacci(i))
	}
	fmt.Println(position)
}

func fibonacci(position uint) uint {
	if position <= 1 {
		return position
	}

	return fibonacci(position-2) + fibonacci(position-1)
}
