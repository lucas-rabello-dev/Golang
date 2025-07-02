package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"fmt"
)

type Product struct {
	ID int `gorm: primaryKey`
	Name string
	Price float64
	gorm.Model
}

func main() {
	db, err := gorm.Open(sqlite.Open("teste.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}

	db.AutoMigrate(&Product{})

	defer fmt.Println("Operação concluída!")

	// Inserindo dados
	// products := []Product {
	// 	{Name: "Notebook", Price: 5000.00},
	// 	{Name: "Celular", Price: 1200.00},
	// }

	// db.Create(&products)

	// Atualizando produto
	// var p Product
	// db.First(&p, 1)
	// p.Name = "Teste"
	// db.Save(&p)

	// Excluindo produto
	var p2 Product
	db.First(&p2, 3)
	db.Delete(&p2)

}