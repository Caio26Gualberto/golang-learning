package main

import "fmt"

func main() {
	total := sum(5, 10, 23, 13)
	fmt.Println(total)

	write("Olá mundo!", 1, 2, 3)
}

func sum(n1 ...int) int {
	total := 0
	for _, number := range n1 {
		total += number
	}

	return total
}

func write(text string, numbers ...int) {
	for _, result := range numbers {
		fmt.Println(text, result)
	}
}
