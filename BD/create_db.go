package main

import (
	"database/sql"
	"log"
	_ "github.com/mattn/go-sqlite3"
)


func main() {
	db, err := sql.Open("sqlite3", "./teste.db") // .Open() cria ou abre um arquivo .db
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	// código em SQL
	sqlStmt := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		name TEXT
	);
	`
	// executar o comando em SQL
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Tabela de 'users' foi criada.")


}