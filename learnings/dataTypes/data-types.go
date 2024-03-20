package main

import (
	"errors"
	"fmt"
)

func main() {
	var number int = 10000
	number2 := 10000
	fmt.Println(number)
	fmt.Println(number2)

	//uint is a unsygned int
	//alis
	//INT32 = RUNE

	var number3 rune = 123456
	fmt.Println(number3)

	//BYTE = UINT8
	var number4 byte = 123
	fmt.Println(number4)

	var realNumber float32 = 123.45
	fmt.Println(realNumber)

	var realNumber2 float64 = 1233400.45
	fmt.Println(realNumber2)

	realNumber3 := 12345.67
	fmt.Println(realNumber3)

	var str string = "HAHAHAHAHA"
	fmt.Println(str)

	str2 := "HAHAHAHA 2"
	fmt.Println(str2)

	char := 'B'       // It has to return 66 because it is the position of B in the ASCII table
	fmt.Println(char) // Returns 66 number

	var text string
	fmt.Println(text) // Returns white space non initialized variable

	var numberZero int
	fmt.Println(numberZero) // Returns 0 non initialized variable

	var boolean bool = true
	fmt.Println(boolean)

	boolean2 := false
	fmt.Println(boolean2)

	var error error // Error type by default is null, if you wanted to generate an error, use Go's native package for the same
	fmt.Println(error)

	erro2 := errors.New("Erro interno")
	fmt.Println(erro2)
}
