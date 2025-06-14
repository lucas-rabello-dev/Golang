package main

import (
	"fmt"
	"os"
	"strconv"
	s "strings"
)


func main() {



	// ------------------------------------------ os strconv -----------------------------------------------------------
	// input da compilação 1 (str)
	valor1 := os.Args[1]
	fmt.Println("Primeiro valor digitado:",valor1)
	// input da compilação 2 (str)
	valor2 := os.Args[2]
	fmt.Println("Segundo valor digitado:",valor2)
	// input da compilação 1 de str para (int)
	valor3, _ := strconv.Atoi(os.Args[1])
	fmt.Println("Valor convertido de str para int", valor3)
	// input da compilação 1 de str para (int)
	valor4, _ := strconv.Atoi(os.Args[2])
	fmt.Println("Valor convertido de str para int", valor4)
	soma := valor3 + valor4
	fmt.Printf("A soma dos valores: %d \n", soma)

	// ----------------------------------------------------------- strings------------------------------------------------------

	texto := `
Python é uma boa linguagem de programação para estudar lógica, porém JavaScript abre mais portas para o futuro!`
	palavra := "Python"
	fmt.Printf("Texto: %s \nPalavra: %s \n", texto, palavra)

	// varifica se existe "Python" na váriavel texto e retorna bool
	texto1 := s.Contains(texto, palavra)
	if texto1 == true {
		fmt.Printf("A palavra %s está no texto! \n", palavra)
	} else {
		fmt.Printf("A palavra %s não está no texto! \n", palavra)
	}

	// transforma todas as letras em minúsculo
	fmt.Println("Texto com todas as letras em minúsculo:")
	fmt.Println(s.ToLower(texto))

	fmt.Printf("\n")

	// transforma todas as Letras em maiúsculo
	fmt.Println("Texto com todas as letras em maiúsculo:")
	fmt.Println(s.ToUpper(texto))

	// Arrays tem tamanho obrigatório e fixo
	// Arrays com tamanho definido
	var numeros [5]int

	// dando o valor por indice
	numeros[0] = 123

	fmt.Println(numeros)

	var numeros2 = [5]int{10, 20, 30, 40, 50}
	fmt.Println(numeros2)

	// slice (mais flexiveis sem precisar de tamanhos definidos)

	// slice vazio
	var numeros3 = []int{}

	fmt.Println(numeros3)


}