package main

import (
	adresses "first/Adresses"
	"fmt"
)

func main() {
	adressType := adresses.AdressType("Avenue Paulista")

	fmt.Println(adressType)
}
