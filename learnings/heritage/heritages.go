package main

import "fmt"

type Person struct {
	name    string
	surname string
	age     uint8
	height  float32
}

type Student struct {
	Person
	course     string
	university string
}

func main() {
	p1 := Person{"Caio", "Gualberto", 23, 1.80}
	fmt.Println(p1)

	e1 := Student{p1, "Análise e Desenvolvimento de Sistemas", "Estácio"}
	fmt.Println(e1)
}
