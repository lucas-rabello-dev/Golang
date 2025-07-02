package main


import "fmt"


func main() {
	// slices são iguais as arrays mas não precisa definir um tamanho
	// e o tamanho é flexível

	// com make()
	numeros := make([]int, 5) // cria um slice com 5 elementos zerados

	nomes := []string{"teste1", "teste2", "teste3"}

	numeros = append(numeros, 10, 20, 30, 405, 60)

	fmt.Println(nomes)

	for _, valor := range numeros {
		fmt.Printf("Valor: %d", valor)
	}

	fmt.Println(len(nomes))

	


}