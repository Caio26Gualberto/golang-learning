package main

import "fmt"

func main() {
	// Arithmetics + - / * %

	fmt.Println("----- Arithmetics -----")
	sum := 1 + 2
	subtract := 2 - 2
	division := 10 / 4
	multiplication := 10 * 5
	restOfDivision := 10 % 2

	fmt.Println(sum, subtract, division, multiplication, restOfDivision)

	// var number1 int16 = 10  It's impossible to add this because Go is strongly typed
	// var number2 int32 = 15
	// soma := number1 + number2
	// fmt.Println(soma)

	var number1 int16 = 10
	var number2 int16 = 15
	soma := number1 + number2
	fmt.Println(number1, number2, soma)

	// Atribuition

	fmt.Println("----- Atribuition -----")
	var variable string = "Caio"
	variable2 := "Gualberto"
	fmt.Printf(variable, variable2)

	// Relational Operators

	fmt.Println("----- Relational Operators -----")
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 != 2)

	// Logic Operators

	fmt.Println("----- Logic Operators -----")
	t, f := true, false
	fmt.Println(t && f)
	fmt.Println(t || f)
	fmt.Println(!t)

	// Unary Operators

	fmt.Println("----- Unary Operators -----")
	number := 10
	number++
	number += 10
	number--
	number -= 10
	fmt.Println(number)
	fmt.Println(t || f)
	fmt.Println(!t)

	// Ternary Operators

	fmt.Println("----- Ternary Operators -----") // In Go there is no ternary operator to keep the code easier, just if else
	var txt string
	if number > 5 {
		txt = "maior que 5"
	} else {
		txt = "Menor que 5"
	}
	fmt.Println(txt)

}
