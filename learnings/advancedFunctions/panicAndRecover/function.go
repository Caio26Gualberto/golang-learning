package main

import "fmt"

func main() {
	fmt.Println(studentHasApproved(6, 6))
	fmt.Println("Pós execução !")
}

func studentHasApproved(n1, n2 float32) bool {
	defer recoverExecution()
	average := (n1 + n2) / 2

	if average > 6 {
		return true
	} else if average < 6 {
		return false
	}

	panic("A média é exatamente 6")
}

func recoverExecution() {
	if r := recover(); r != nil {
		fmt.Printf("Execução recuperada com sucesso!")
	}
}
