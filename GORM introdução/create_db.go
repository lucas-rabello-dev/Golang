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
}