package auxiliary

import "fmt"

func Write() {
	fmt.Println("Escrevendo de auxiliar")
	write2() // Can only be called from here because it's from the same auxiliary package
}
