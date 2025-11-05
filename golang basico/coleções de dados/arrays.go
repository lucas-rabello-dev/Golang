package main


import "fmt"


func main() {
	// declaração com var com 5 de tamanho
	var numeros [5]int

	// atribuir um valor
	numeros[0] = 10
	numeros[0] = 20

	var nomes = [2]string{"teste1", "teste2"}

	// declaração sem var com 5 de tamanho
	numeros2 := [5]int{}

	numeros2[0] = 90

	// ... fala ao compilador definir o tamanho 
	nomes2 := [...]string{"teste1", "teste2"}


	// acessar elementos um por um 
	fmt.Println(nomes[0])

	// mostrar elementos usando o for 
	for i := 0; i < len(nomes2); i++ {
		fmt.Println(nomes2[i])
	}

	// usando range (foreach)
	for i, v := range nomes2 {
		fmt.Printf("Indice: %d Valor: %s \n", i, v)
	}

}