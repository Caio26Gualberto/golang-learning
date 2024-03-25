package main

import "fmt"

type User struct {
	name string
	age  uint8
}

func main() {
	user := User{"Caio", 22}
	user.save()
	user.userIsLegalAge()

	user2 := User{"Lorenzo", 20}
	user2.save()
	user2.userIsLegalAge()
	user2.happyBirthday()
	fmt.Println(user2.age)
}

func (u User) save() {
	fmt.Printf("Salvando os dados do usuário %s no banco de dados\n", u.name)
}

func (u User) userIsLegalAge() bool {
	if u.age >= 18 {
		fmt.Printf("Usuário %s é maior de idade tem %d anos\n", u.name, u.age)
		return true
	}

	fmt.Printf("Usuário %s é menor de idade tem %d anos\n", u.name, u.age)
	return false
}

func (u *User) happyBirthday() {
	u.age++
}
