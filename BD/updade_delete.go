package main


import (
	"database/sql"
	"log"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./teste.db")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	// Atualizando dados
	result, err := db.Exec("UPDATE users SET name = ? WHERE id = ?", "Teste", 1) // muda o item do ID i para Teste
	if err != nil {
		log.Fatal(err)
	}
	// verifica quantas linhas foram afetadas
	affectedRows, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Updated %d row(s)", affectedRows)

}
