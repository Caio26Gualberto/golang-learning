package main

import "fmt"

func main() {
	text := "Dentro da função main"
	fmt.Println(text)

	newFunction := closure()
	newFunction()
}

func closure() func() {
	text := "Dentro da função closure"

	function := func() {
		fmt.Println(text)
	}

	return function
}
