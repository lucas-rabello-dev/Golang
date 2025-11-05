package main

import (
	"fmt"
)


func main() {

	notas := []float64{9.8, 7.3, 6.8, 2.2, 9.9}

	var total float64

	for _, nota := range notas {
		total += nota
	}

	media := float64(len(notas))

	fmt.Println("Média das notas: ", media)




}

