package main

import "fmt"

func main() {
	function1()
	function2()

	// defer function1()
	function2()

	fmt.Println(studentHasApproved(5, 7))
}

func function1() {
	fmt.Println("Executando função 1")
}

func function2() {
	fmt.Println("Executando função 2")
}

func studentHasApproved(n1, n2 float32) bool {

	defer fmt.Println("Média sendo calculada")
	average := (n1 + n2) / 2

	if average >= 6 {
		return true
	}

	return false
}
