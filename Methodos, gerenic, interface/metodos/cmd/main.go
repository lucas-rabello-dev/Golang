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

	fmt.Println()
	// fmt.Println(estoque.ViewLogs())
	logs := estoque.ViewLogs()
	for _, log := range logs {
		fmt.Printf("\n[%s] Ação: %s - Usuário: %s - Item ID: %d - Quantidade: %d - Motivo: %s",
	log.TimeStamp.Format("02/01 15:04:05"), log.Action, log.User, log.ItemId, log.Quantity, log.Reason)
	}
	// log.TimeStamp.Format("02/01 15:04:05") -> para formatar a data corretamente

	fmt.Println("\nO valor total: ", estoque.CalculateTotalCost())

	itemByName, err := services.FindByName(itens, "Teste 1")
	if err != nil {
		panic(err)
	}
	fmt.Println(itemByName)

	searchItem, err := services.FindBy(itens, func(item models.Item) bool {
		// return item.Name == "Teste 2"
		return item.Price <= 300.00
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(searchItem)




}
