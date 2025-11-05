package main

import (
	"fmt"
)


func main() {
	
	estudantes := map[string]float64{
		"Eduardo": 7.5,
		"Lucas": 9.0,
		"Pedro": 8.7,
		"Leonardo": 5.8,
	}

	for nome, nota := range estudantes {

		status := "Reprovado"

		if nota >= 6 {
			status = "Aprovado"
		}

		fmt.Printf("Nome: %s Nota: %.0f Status: %s \n", nome, nota, status)
	}
	


}
