package main

import "fmt"

func main() {
	var variable1 int = 10
	var variable2 int = variable1

	fmt.Println(variable2)

	variable1++
	fmt.Println(variable2)

	var variable3 int
	var pointer *int // The * is used to say that I am expecting a type "x" with the memory address, that is, here I am expecting a memory address that has the type int stored

	variable3 = 100
	pointer = &variable3 // With & I'm saying that I need the memory address of this pointer

	fmt.Println(variable3, *pointer) // Using * again I am saying that I need the value stored in this pointer, if you remove the * it will return the address where the value is stored

}
