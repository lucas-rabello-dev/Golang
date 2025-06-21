package main 

import (
	"log"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/google/uuid"
	"fmt"
)

type User struct {
	ID string
	Name string
	Email string
	Age int
}

// retorna os valores diretamente para a struct 
func NewUser(name string, email string, age int) *User{ // retorna um ponteiro para a struct garantindo que mude seu valor sem ser uma cópia
	return &User { // aponta para o endereço para garantir que o valor mude 
		ID: uuid.New().String(),
		Name: name,
		Email: email,
		Age: age,
	}
}

func createTable(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS users(
		id TEXT PRIMARY KEY,
		name TEXT,
		email TEXT,
		age INTEGER	
	);
	`
	_, err := db.Exec(query)

	return err
}

func insertUser(db *sql.DB, user *User) error {
	stmt, err := db.Prepare("INSERT INTO users(id, name, email, age) VALUES($1, $2, $3, $4)")
	if err != nil {
		return err 
	}

	defer stmt.Close()

	_, err = stmt.Exec(user.ID, user.Name, user.Email, user.Age)
	if err != nil {
		return err
	}
	return nil
}

func updadeUser(db *sql.DB, user *User) error {
	stmt, err := db.Prepare("UPDATE users SET name = ?, email = ?, age = ? WHERE id = ?")
	if err != nil {
		return err 
	}

	defer stmt.Close()

	_, err = stmt.Exec(user.Name, user.Email, user.Age, user.ID)
	if err != nil {
		return err
	}
	return nil
}

func selectUsers(db *sql.DB) ([]User, error) {
	rows, err := db.Query("SELECT id, name, email, age FROM users")
	if err != nil {
		return nil, err
	}
 
	defer rows.Close()
	var users []User

	for rows.Next() {
		var u User
		err = rows.Scan(&u.ID, &u.Name, &u.Email, &u.Age)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}


func deleteUsers(db *sql.DB, id string) error {
	stmt, err := db.Prepare("DELETE FROM users WHERE id = ?")
	if err != nil {
		return err 
	}

	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return nil
}



func main(){
	// cira ou abre o DB
	db, err := sql.Open("sqlite3", "./app.db")
	if err != nil {
		log.Fatal(err)
	}
	// fecha ao final do programa
	defer db.Close()
	// Criando a tabela
	err = createTable(db)
	if err != nil {
		log.Fatal(err)
	}

	// inserindo usuario
	user := NewUser("Lucas", "lucas@exemplo.com", 21)
	err = insertUser(db, user)
	if err != nil {
		log.Fatal(err)
	}

	// atualização de usuário
	user.Name = "teste"
	user.Age = 33
	err = updadeUser(db, user)
	if err != nil {
		log.Fatal(err)
	}

	// listagem de usuários
	users, err := selectUsers(db)
	if err != nil {
		log.Fatal(err)
	}

	for _, u := range users {
		fmt.Printf("Usuário %s possui o email %s e possui idade %d anos\n", u.Name, u.Email, u.Age)
	}

	// exclusão de dados
	err = deleteUsers(db, user.ID)
	if err != nil {
		log.Fatal(err)
	}

}