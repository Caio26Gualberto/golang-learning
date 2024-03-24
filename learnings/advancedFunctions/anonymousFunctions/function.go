package main

import "fmt"

func main() {
	func() {
		fmt.Println("Ola mundo")
	}()

	func(text string) {
		fmt.Println(text)
	}("Olá mundo")

	returnn := func(text string) string {
		return fmt.Sprintf("Recebendo -> %s", text)
	}("Olá mundo")

	fmt.Println(returnn)
}
