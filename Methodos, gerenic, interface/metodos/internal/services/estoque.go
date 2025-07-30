package services

import "estoque/internal/models"

type Estoque struct {
	itens map[string]models.Item
	logs []models.Log
}


// Passar informações do estoque por valor
func NewEstoque() *Estoque {
	return &Estoque{
		itens: make(map[string]models.Item),
		logs: []models.Log{},
	}
}