package main

//To declare a private function the first caracter must be lower, and if you want a public function declare with first caracter Upper
// Cmd command go tidy remove all dependencies with no usage in your application

import (
	"fmt"
	"modules/auxiliary"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Olá mundo!")
	auxiliary.Write()
	error := checkmail.ValidateFormat("caiogualbertodev@outlook.com")
	fmt.Println(error)
}
