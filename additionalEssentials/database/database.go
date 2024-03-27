package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	connectioString := "golang:golang@/golang?charset=utf8&parseTime=True&loc=Local"
	db, erro := sql.Open("mysql", connectioString)

	if erro != nil {
		log.Fatal(erro)
	}
	defer db.Close()

	if erro = db.Ping(); erro != nil {
		log.Fatal(erro)
	}

	fmt.Println("Conexão está aberta")

	users, erro := db.Query("select * from users")
	defer users.Close()

	fmt.Println(users)

}
