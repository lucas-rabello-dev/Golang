// Relacionamento nesse caso é como se fosse uma tabela cursos com muitos usuarios
// e usuarios com muitos cursos
// parecido com has many
package main

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Category struct {
	ID       int `gorm:"primerykey"`
	Name     string
	Products []Product `gorm:"many2many:products_categories;"`
}

type Product struct {
	ID           int `gorm: primaryKey`
	Name         string
	Price        float64
	Categories     []Category `gorm:"many2many:products_categories;"`
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

	// category2 := Category{Name: "Eletrodomésticos"}
	// db.Create(&category2)

	// inseindo uma categoria
	// db.Create(&Product{
	// 	Name:       "Microondas",
	// 	Price:      679.00,
	// 	Categories: []Category{category, category2},
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