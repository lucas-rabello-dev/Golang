package main

// o uso de _ é para dizer que não vão ser usadas as funções do pacote
// no caso desse código essa lib é usada para os registros no driver
import (
	"database/sql"
	"log"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./teste.db") // ./ para indicar que go deve procurar ou criar o arquivo no diretorio atual
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	linhas, err := db.Query("SELECT id, name FROM users")
	if err != nil {
		log.Fatal(err)
	}
	
	defer linhas.Close()

	for linhas.Next() {
		var id int
		var name string
		err := linhas.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("ID: %d   Nome: %s", id, name)
	}

	if err = linhas.Err(); err != nil {
		log.Fatal(err)
	}


	
}