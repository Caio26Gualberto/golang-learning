package main

import "fmt"

type User struct {
	name   string
	age    uint8
	adress Adress
}

type Adress struct {
	street string
	number uint8
}

func main() {
	var u User
	var a1 Adress

	u.name = "Caio"
	u.age = 22

	a1.number = 30
	a1.street = "Long Avenue"
	u.adress = a1

	fmt.Println(u)

	u2 := User{"Caio", 23, Adress{street: "Hollywood", number: 10}} // Same as constructor
	fmt.Println(u2)

	u3 := User{name: "Caio"}
	fmt.Println(u3)
}
