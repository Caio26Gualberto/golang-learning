package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type Cachorro struct {
	Name string `json:"name"`
	Race string `json:"race"`
	Age  uint   `json:"age"`
}

func main() {
	c := Cachorro{"Ruan", "Lhasa apso", 5}
	fmt.Println(c)

	cInJson, erro := json.Marshal(c)
	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(cInJson)
	fmt.Println(bytes.NewBuffer(cInJson))

	c2 := map[string]string{
		"name": "Toby",
		"race": "Poodle",
	}

	cInJson2, erro2 := json.Marshal(c2)
	if erro2 != nil {
		log.Fatal(erro2)
	}
	fmt.Println(bytes.NewBuffer(cInJson2))
}
