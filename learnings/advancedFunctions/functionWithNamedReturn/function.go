package main

import "fmt"

func main() {
	sum, subtract := calc(10, 20)
	fmt.Println(sum, subtract)
}

func calc(n1, n2 int) (sum int, subtract int) {
	sum = n1 + n2
	subtract = n1 - n2
	return
}
