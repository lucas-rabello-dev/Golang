package main

import "fmt"


func fullName(firstName, lastName string) {
	fmt.Printf("Nome completo: %s %s \n", firstName, lastName)

}

func sum(num1, num2 int) {
	soma := num1 + num2
	fmt.Printf("A soma de %d + %d é %d \n", num1, num2, soma)
}


func main() {
	fmt.Println("Utilizando função com parametros")
	fullName("Lucas", "Carvalho")
	sum(12, 22)
}