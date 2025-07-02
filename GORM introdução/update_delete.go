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

	// Atualizando um produto
	// var p Product
	// db.First(&p, 1) // busca produto numero 1 -> ID == 1
	// p.Name = "Atualizando dados"
	// db.Save(&p)

	// Atualizando multiplas colunas | peecisa especificar o ID
	db.Model(&Product{}).Where("id = ?", 1).Updates(Product{Price: 4930, Name: "Carrinho de brinquedo"})

	// Exclusão de dados
	var p3 Product
	db.First(&p3, 4) // seleciona por ID
	db.Delete(&p3) // Deleta o item com ID 4


}