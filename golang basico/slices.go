package main

import (
	"fmt"
)


func main() {

	
	frutas := []string{"fruta1", "fruta2", "fruta3"}

	// retorna o item do incice 1 até o 3 
	subslice := frutas[1:3]

	fmt.Println(frutas, subslice)



}