package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0

	for i <= 10 {
		time.Sleep(time.Second)
		fmt.Println(i)
		i++
	}

	for j := 0; j <= 10; j++ {
		fmt.Println(j)
	}

	nomes := [3]string{"Caio", "Lorenzo", "João"}

	for _, nome := range nomes {
		fmt.Println(nome)
	}

	for _, letra := range "Caio" {
		fmt.Println(string(letra))
	}

	user := map[string]string{
		"nome":      "Caio",
		"sobrenome": "Gualberto",
	}

	fmt.Println(user)

	for _, valor := range user {
		fmt.Println(valor)
	}

	type User struct {
		nome      string
		sobrenome string
	}

	// user2 := User{"Caio", "Gualberto"}

	// for _, valor := range user2 { Range dont work in structs
	// 	fmt.Printf(valor)
	// }
}
