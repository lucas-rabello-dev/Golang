package main

import (
	"estoque/internal/models"
	"fmt"
)


func main() {
	item1 := models.Item{
		ID: 1,
		Name: "Ps5",
		Quantity: 2,
		Price: 1200.00,
	}

	fmt.Println(item1.Info())

	item2 := models.Item{
		ID: 2,
		Name: "Ps4",
		Quantity: 4,
		Price: 800.00,
	}

	fmt.Println(item2.Info())

	
}