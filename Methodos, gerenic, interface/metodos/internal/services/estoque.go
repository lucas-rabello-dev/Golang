package services

import (
	"estoque/internal/models"
	"fmt"
	"strconv"
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