package main

import "fmt"

func main() {
	number1 := 20
	numberWithInvertedSignal := signalInvert(number1)
	fmt.Println(numberWithInvertedSignal)
	fmt.Println(number1)

	signalInvertWithPointer(&number1)
	fmt.Println(number1)
}

func signalInvert(n1 int) int {
	return n1 * -1
}

func signalInvertWithPointer(n1 *int) {
	*n1 = *n1 * -1
}
