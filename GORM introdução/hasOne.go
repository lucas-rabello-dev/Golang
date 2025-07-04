package main

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type SerialNumber struct {
	ID int `gorm: "primaryKey"`
	Number string
	ProductID int
}

type Category struct {
	ID   int `gorm: "primerykey"`
	Name string
}

type Product struct {
	ID         int `gorm: primaryKey`
	Name       string
	Price      float64
	CategoryId int // armazenar o ID
	Category   Category
	SerialNumber SerialNumber
	gorm.Model
}

func main() {
	db, err := gorm.Open(sqlite.Open("teste.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}

	db.AutoMigrate(&Product{}, &Category{}, SerialNumber{})

	defer fmt.Println("Operação concluída!")

	// Inserindo a categoria
	category := Category{Name: "Cozinha"}
	db.Create(&category)
	
	// inseindo um produto
	db.Create(&Product{
		Name: "Fogão",
		Price: 3600.00,
		CategoryId: category.ID,
	})

	// inserindo Serial Number
	db.Create(&SerialNumber{
		Number: "12323",
		ProductID: 1,
	})

	// Listando produtos
	var products []Product
	db.Preload("Category").Preload("Serial Number").Find(&products)
	for _, product := range products {
		fmt.Println(product.Name, product.Category.Name, product.Price, product.SerialNumber.Number)
	}
}
