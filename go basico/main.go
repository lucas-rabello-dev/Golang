package main

import (
	"fmt"
	
)

func main() {
		
	var (
		operacao string
		calculo int
		num1 int
		num2 int
	)


	fmt.Print("Digite a operação: ")
	fmt.Scan(&operacao)

	if operacao != "+" && operacao != "-" && operacao != "*" && operacao != "/"{
		fmt.Println("Digite uma opção válida!")
		return
	}



	fmt.Print("Digite o primeiro número: ")
	fmt.Scan(&num1)

	fmt.Print("Digite o segundo número: ")
	fmt.Scan(&num2)

	switch operacao{
	case "+":
		calculo = num1 + num2
	case "-":
		calculo = num1 - num2

	case "*":
		calculo = num1 * num2
	case "/":
		calculo = num1 / num2
	}
	
	fmt.Println("O resultado é:", calculo)



}


