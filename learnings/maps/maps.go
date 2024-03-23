package main

import (
	"fmt"
)

func main() {

	user := map[string]string{
		"nome":      "Caio",
		"sobrenome": "Gualberto",
	}

	fmt.Println(user)
	fmt.Println(user["nome"])

	user2 := map[string]map[string]string{
		"nome": {
			"primeiro": "Caio",
			"segundo":  "Gualberto",
		},
		"curso": {
			"nome":     "Analise e Desenvolvimento de Sistemas",
			"terminou": "sim",
		},
	}

	fmt.Println(user2)
	delete(user2, "curso")
	fmt.Println(user2)

	user2["Skills"] = map[string]string{
		"Golang": "true",
	}
	fmt.Println(user2)
}
