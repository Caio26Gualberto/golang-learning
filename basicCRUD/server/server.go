package server

import (
	"crud/db"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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
		fmt.Println(erro)
		w.Write([]byte("Erro ao conectar ao banco de dados"))
		return
	}
	defer db.Close()

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

func GetUsers(w http.ResponseWriter, r *http.Request) {
	db, err := db.Connect()

	if err != nil {
		w.Write([]byte("Erro ao conectar com o banco de dados!"))
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM users")

	if err != nil {
		w.Write([]byte("Erro ao bsucar os usuários!"))
		return
	}
	defer rows.Close()

	var users []user
	for rows.Next() {
		var user user

		if err := rows.Scan(&user.Id, &user.Name, &user.Email); err != nil {
			w.Write([]byte("Erro ao escanear o usuário"))
			return
		}

		users = append(users, user)
		_ = users
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(users); err != nil {
		w.Write([]byte("Erro ao converter usuários para JSON"))
		return
	}
}

func SearchUser(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	id, err := strconv.ParseUint(parameters["id"], 10, 32)

	if err != nil {
		w.Write([]byte("Erro ao converter o parâmetro para inteiro"))
		return
	}

	db, err := db.Connect()

	if err != nil {
		w.Write([]byte("Erro ao conectar com o banco"))
		return
	}

	row, err := db.Query("SELECT * FROM users WHERE id = ?", id)

	if err != nil {
		w.Write([]byte("Erro ao recuperar usuário do banco"))
		return
	}

	var user user
	if row.Next() {
		if err := row.Scan(&user.Id, &user.Name, &user.Email); err != nil {
			w.Write([]byte("Erro ao escanear o usuário do banco"))
			return
		}
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(user); err != nil {
		w.Write([]byte("Erro ao converter o usuário para JSON"))
		return
	}

}
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	id, err := strconv.ParseUint(parameters["id"], 10, 32)

	if err != nil {
		w.Write([]byte("Erro ao converter o parâmetro da requisição para inteiro"))
		return
	}

	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("erro ao ler o corpo da requisição"))
		return
	}

	var user user
	if err := json.Unmarshal(requestBody, &user); err != nil {
		w.Write([]byte("erro ao converter o corpo da requisição para JSON"))
		return
	}

	db, err := db.Connect()
	if err != nil {
		w.Write([]byte("erro ao converter o corpo da requisição para JSON"))
		return
	}
	defer db.Close()

	statement, err := db.Prepare("UPDATE users SET name = ?, email = ? WHERE id = ?")
	if err != nil {
		w.Write([]byte("erro ao atualizar statement"))
		return
	}
	defer statement.Close()

	if _, err := statement.Exec(user.Name, user.Email, id); err != nil {
		w.Write([]byte("Erro ao atualizar usuário"))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	id, err := strconv.ParseUint(parameters["id"], 10, 32)

	if err != nil {
		w.Write([]byte("Erro ao converter o parâmetro para inteiro"))
		return
	}

	db, err := db.Connect()
	if err != nil {
		w.Write([]byte("Erro ao conectar com o banco de dados"))
		return
	}
	defer db.Close()

	statement, err := db.Prepare("DELETE FROM users WHERE id = ?")
	if err != nil {
		w.Write([]byte("Erro ao criar o statement"))
		return
	}
	defer statement.Close()

	if _, err := statement.Exec(id); err != nil {
		w.Write([]byte("Erro ao deletar o usuário"))
		return
	}

	w.WriteHeader(http.StatusNoContent)

}
