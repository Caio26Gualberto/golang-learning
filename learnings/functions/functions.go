package main

import "fmt"

func main() {
	sum := sum(5, 5)
	fmt.Println(sum)

	f("Função")
	sumResult, subtractResult := mathCalcs(5, 5)

	sumResult2, _ := mathCalcs(10, 10)

	fmt.Println(sumResult, subtractResult)

	fmt.Println(sumResult2)
}

func sum(number1 int8, number2 int8) int8 {
	return number1 + number2
}

var f = func(txt string) {
	fmt.Println(txt)
}

func mathCalcs(number1, number2 int8) (int8, int8) {
	sum := number1 + number2
	subtraction := number1 - number2

	return sum, subtraction
}
