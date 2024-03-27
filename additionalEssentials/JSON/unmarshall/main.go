package main

import (
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
	cInJson := `{"name": "Ruan","race": "Lhasa apso","age": 5}`

	var c Cachorro

	if erro := json.Unmarshal([]byte(cInJson), &c); erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(c)

	cInJson2 := `{"name": "Tody","race": "Poodle"}`

	c2 := make(map[string]string)

	if erro2 := json.Unmarshal([]byte(cInJson2), &c2); erro2 != nil {
		log.Fatal(erro2)
	}

	fmt.Println(c2)

}
