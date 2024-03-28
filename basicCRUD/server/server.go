package server

import (
	"crud/db"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type user struct {
	Id    uint32 `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	requestBody, erro := io.ReadAll(r.Body)

	if erro != nil {
		w.Write([]byte("Falha ao ler o corpo da requisição"))
		return
	}

	var user user

	if erro = json.Unmarshal(requestBody, &user); erro != nil {
		w.Write([]byte("Erro ao converter o usuário para struct"))
		return
	}

	fmt.Println(user)

	db, erro := db.Connect()

	if erro != nil {
		w.Write([]byte("Erro ao conectar ao banco de dados"))
		return
	}

	// PREPARE STATEMENT
	statement, erro := db.Prepare("INSERT INTO users (name, email) values (?, ?)")

	if erro != nil {
		w.Write([]byte("Erro ao criar o statement"))
		return
	}
	defer statement.Close()

	insert, erro := statement.Exec(user.Name, user.Email)

	if erro != nil {
		w.Write([]byte("Erro ao executar o statement"))
		return
	}

	insertedId, erro := insert.LastInsertId()
	if erro != nil {
		w.Write([]byte("Erro ao obter o Id inserido"))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Usuário inserido com sucesso! id: %d", insertedId)))
}
