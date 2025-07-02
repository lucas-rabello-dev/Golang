package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"fmt"
)

type Category struct {
	ID int `gorm: "primerykey"`
	Name string
}

type Product struct {
	ID int `gorm: primaryKey`
	Name string
	Price float64
	CategoryId int // armazenar o ID
	Category Category
	gorm.Model
}

func main() {
	db, err := gorm.Open(sqlite.Open("teste.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}

	db.AutoMigrate(&Product{}, &Category{})

	defer fmt.Println("Operação concluída!")

	// Inserindo a categoria
	// category := Category{Name: "Cozinha"}
	// db.Create(&category)
	
	// // inseindo um produto
	// db.Create(&Product{
	// 	Name: "Fogão",
	// 	Price: 3600.00,
	// 	CategoryId: category.ID,
	// })

	// Listando produtos

	var products []Product
	db.Preload("Category").Find(&products)
	for _, product := range products {
		fmt.Println(product.Name, product.Category.Name, product.Price)
	}

}