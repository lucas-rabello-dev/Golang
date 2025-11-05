package tax



import (
	"testing"
)

func TestCalculateTax(t *testing.T) {
	amount := 500.00
	expected := 5.0
	result := CalculateTax(amount)

	if result != expected {
		t.Errorf("Expected %f but got %f \n", expected, result)
	}
}


// teste em lote
func TestCAlculate(t *testing.T) {
	type calcTax struct {
		amount, expected float64
	}

	table := []calcTax{
		{500.00, 5.0},
		{1000.00, 10.0},
		{1500.00, 10.0},
	}

	for _, item := range table {
		result := CalculateTax(item.amount)
		if result != item.expected {
			t.Errorf("Expected %f but got %f", item.expected, result)

		}
	}
}

// para executar
// go test -bench=<nome da função>

// benchmak mede o desenpenho das funções
// usar -json no final do comando go para retornar dados de json
// -v também para mais informações
// -cout=5 para excutar 5 vezes

func BenchmarkCalculateTax(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateTax(500.0)
	}
}

func BenchmarkCalculateTax2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateTax2(500.0)
	}
}

// para rodar
// go test -fuzz=<nome da função>
func FuzzCalculateTax(f *testing.F) {
	seed := []float64{1, 2, 3, 500.0, 1000.0, 1500.0}
	for _, amount := range seed {
		f.Add(amount)
	}

	f.Fuzz(func(t *testing.T, amount float64) {
		result := CalculateTax(amount)
		if amount <= 0 && result != 0 {
			t.Errorf("Recived %f but expected 0", result)
		}
	})
}