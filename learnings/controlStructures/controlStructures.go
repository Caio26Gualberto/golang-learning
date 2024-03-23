package main

import "fmt"

func main() {

	number := 10

	if number > 15 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("Menor que 15")
	}

	if number2 := 15; number2 > 0 {
		fmt.Println("Number2 é maior que 0")
	} else if number2 < -10 {
		fmt.Println("number2 é menor que -10")
	} else {
		fmt.Println("number2 é menor que 0")
	}
}
