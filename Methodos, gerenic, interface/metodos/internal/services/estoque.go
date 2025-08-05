package services

import (
	"estoque/internal/models"
	"fmt"
	"strconv"
	"time"
)

type Estoque struct {
	itens map[string]models.Item
	logs []models.Log
}


// Passar informações do estoque por referência
func NewEstoque() *Estoque {
	return &Estoque{
		itens: make(map[string]models.Item),
		logs: []models.Log{},
	}
}


// Método para adicionar Estoque
func (e *Estoque) AddItem(item models.Item) error {
	if item.Quantity <= 0 {
		return fmt.Errorf("erro ao adicionar o produto, a quantidade não pode ser menor ou igual a zero Qtd atual: %d ", item.Quantity)
	}
	existingItem, exists := e.itens[strconv.Itoa(item.ID)]
	if exists {
		item.Quantity += existingItem.Quantity
	}
	e.itens[strconv.Itoa(item.ID)] = item

	// Parte dos logs
	e.logs = append(e.logs, models.Log{
		TimeStamp: time.Now(),
		Action: "Entrada de Item no Estoque",
		User: "Administrador",
		ItemId: item.ID,
		Quantity: item.Quantity,
		Reason: "Adicionando novos itens ao estoque",
	})
	return nil
}

// Método para Listar Itens do Estoque
func (e *Estoque) ListItens() []models.Item {
	var itensList []models.Item // Lista dos itens (slice) de (maps) que vamos receber na main
	for _, item := range e.itens {
		itensList = append(itensList, item)
	}
	return itensList
}


// Método para listar logs
func (e *Estoque) ViewLogs() []models.Log {
	return e.logs
}

// Método para calcular o preço total
func (e *Estoque) CalculateTotalCost() float64 {
	var totalCost float64
	for _, item := range e.itens {
		totalCost += float64(item.Quantity) * item.Price
	}
	return totalCost
}

func FindByName(data []models.Item, name string) ([]models.Item, error){
	var result []models.Item
	for _, item := range data {
		if item.Name == name {
			result = append(result, item)
		}
	}

	if len(result) == 0 {
		return nil, fmt.Errorf("nenhum item com o nome '%s' encontrado", name)
	}
	return result, nil
}


// Usando Generics
func FindBy[T any] (data []T, comparator func(T) bool) ([]T, error){
	var result []T
	for _, v := range data {
		if comparator(v) {
			result = append(result, v)
		}
	}
	if len(result) == 0 {
		return nil, fmt.Errorf("nenhum item foi encontrado")
	}
	return result, nil
}