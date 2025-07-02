package main

import (
	"fmt"

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

	// // Selecionando 1 Produto
	// var product Product
	// db.First(&product, 2) // Busca o produto com base no ID, no caso 2 
	// fmt.Println(product.ID, product.Name, product.Price)

	// var product2 Product
	// // Selecionar por coluna
	// db.First(&product2, "name = ?", "Celular") // busca o produto "Celular"
	// fmt.Println(product2.ID, product2.Name, product2.Price)


	// // Selecionar todos os produtos
	// var products []Product
	// // db.Find(&products) // mostra todos
	// db.Limit(2).Find(&products) // mostra os 2 primeiros no .db
	// for _, product := range products {
	// 	fmt.Println(product)
	// }


	// Consulta com WHERE 
	var products []Product
	// procura por price
	db.Where("price > ?", 3500).Find(&products) // apenas produtos com o preço maior que 3500
	// procura produtos que começam ou terminal com %item% -> item pode ser qualquer palavra
	db.Where("name LIKE ?", "%book%").Find(&products)
	if len(products) == 0 {
		fmt.Println("Nenhum produto encontrado")
		return
	}
	for _, product := range products {
		fmt.Println(product)
	}
	
}