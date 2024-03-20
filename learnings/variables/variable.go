package main

import "fmt"

func main() {
	var variable string = "Caio"
	secondVariable := "Gualberto"

	fmt.Println(variable)
	fmt.Println(secondVariable)

	var (
		thirdVariable  string = "Enjoy"
		fourthVariable string = "Development"
	)

	fmt.Println(thirdVariable, fourthVariable)

	fifthVariable, sixthVariable := "and like", "Doritos"

	fmt.Println(fifthVariable, sixthVariable)

	const constant string = "hahaha"

	fifthVariable, sixthVariable = sixthVariable, fifthVariable

	fmt.Println(fifthVariable, sixthVariable)
}
