package main



import (
	"fmt"
	"math"
)

// acessar o valor de PI
func acessarPI() {
	fmt.Printf("PI Arredondado em duas casa decimais: %.2f \n", math.Pi)
}

// arredondar numeros
func arredondar() {
	num := 12.33
	fmt.Printf("Numero: %d Arrendondado: %d", num, math.Ceil(num)) // Ceil() para cima e Floor() para baixo
}


func main() {
	fmt.Println("Potência de 5 elevado a 5:", math.Pow(5, 5))
	fmt.Println("raiz quadrada de 144:", math.Sqrt(144))
	fmt.Println("Log de 10:", math.Log(10))




	
}