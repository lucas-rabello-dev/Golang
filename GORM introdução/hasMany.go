// mostra os itens que estão dentro de cada categoria
// Relacionamento nesse caso é como se tivesse uma tabela curso com muitos alunos 
package main

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Category struct {
	ID       int `gorm: "primerykey"`
	Name     string
	Products []Product
}

type Product struct {
	ID           int `gorm: primaryKey`
	Name         string
	Price        float64
	CategoryId   int // armazenar o ID
	Category     Category
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

	// inseindo um produto
	// db.Create(&Product{
	// 	Name:       "Microondas",
	// 	Price:      679.00,
	// 	CategoryId: 1,
	// })

	// Listando Categorias e Produtos
	var categories []Category
	err = db.Model(&Category{}).Preload("Products").Find(&categories).Error
	if err != nil {
		panic(err)
	}
	for _, category := range categories {
		fmt.Println(category.Name,":")
		for _, product := range category.Products {
			fmt.Println("->", product.Name, product.Price)
		}
	}

}
