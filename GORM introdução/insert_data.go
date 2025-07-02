package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Product struct {
	ID int `gorm: primaryKey`
	Name string
	Price float64
}

func main() {
	db, err := gorm.Open(sqlite.Open("teste.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}

	// Migrar os dodos da struct para a tabela
	db.AutoMigrate(&Product{})

	// inserindo dados
	db.Create(&Product{Name: "Teste", Price: 3900.00})

	// Outra forma é criando um slice com cada item e seus valores e passar 
	// os dados todos de uma vez
	products := []Product {
		{Name: "Notebook", Price: 2222.00},
		{Name: "Celular", Price: 3200.99},
		{Name: "Tablet", Price: 2500.00},
	}
	db.Create(&products)

	



}