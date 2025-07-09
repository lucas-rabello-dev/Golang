package tax

import (
	"testing"
	"github.com/stretchr/testify/assert"
)


// verifica se é igual uma com a outra
// para rodar
func TestCalculateTax(t *testing.T) {
	// teste para o valor 1000
	tax := CalculateTax(1000.0)
	assert.Equal(t, 10.0, tax, "o imposto para 1000 deve ser 10")

	// teste para o valor 0
	tax = CalculateTax(0)
	assert.Equal(t, 0.0, tax, "O imposto para zero deve ser zero")

	// Teste para valor 500
	tax = CalculateTax(500)
	assert.Equal(t, 5.0, tax, "O imposto para 500 deve ser 5")
}

