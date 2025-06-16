package main

import (
	"fmt"
)

type Carro struct {
	Modelo string
	Ano int
	Cor string
}

func main() {
	carro1 := Carro {
		Modelo: "Fusca",
		Ano: 1965,
		Cor: "Azul",
	}
	
	fmt.Println("Informações do Carro")
	fmt.Printf("Modelo: %s \n",carro1.Modelo)
	fmt.Printf("Ano: %d \n", carro1.Ano)
	fmt.Printf("Cor: %s", carro1.Cor)

}