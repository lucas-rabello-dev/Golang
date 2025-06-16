package main

import "fmt"

// tabuada
func main() {

	var numero int

	fmt.Print("Digite um numero para ver sua tabuada do 1 ao 10: ")
	fmt.Scan(&numero)

	for i := 1; i <= 10; i++ {
		resultado := numero*i
		fmt.Printf("%d x %d = %d \n", numero, i, resultado)
	}



}