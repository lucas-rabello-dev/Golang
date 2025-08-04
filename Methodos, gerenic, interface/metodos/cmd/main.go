package main

import (
	"estoque/internal/models"
	"estoque/internal/services"
	"fmt"
)


func main() {
	// Slice de itens da struct Item
	itens := []models.Item {
		{ID: 1, Name: "Teste 1", Quantity: 8, Price: 122.00},
		{ID: 2, Name: "Teste 2", Quantity: 6, Price: 390.99},
		{ID: 3, Name: "Teste 3", Quantity: 2, Price: 149.76},
	}

	// Faz referência a struct Estoque
	estoque := services.NewEstoque()

	// Percorrer com o index para adicionar todos os itens do Slice 'itens'
	// o segundo valor de range seria um models.Item que é uma struct
	for index := range itens {
		estoque.AddItem(itens[index])
	}

	// Ou pode fazer dessa forma (a do curso)
	// for _, item := range itens {
	// 	err := estoque.AddItem(item)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		continue
	// 	}
	// }

	// fmt.Println(estoque.ListItens())
	for _, item := range estoque.ListItens() {
		fmt.Printf("\n ID: %d | Item: %s | Quantidade: %d| Preço: %.2f",
	item.ID, item.Name, item.Quantity, item.Price)
	}
}
