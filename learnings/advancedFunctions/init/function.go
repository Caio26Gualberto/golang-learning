package main

import "fmt"

func init() {
	fmt.Println("Função init sempre vem primeiro")
}

func main() {
	fmt.Println("Função main sempre vem em segundo")
}
