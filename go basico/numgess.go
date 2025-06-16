package main


import (
	"fmt"
	"math/rand"
	
)


func main() {
	// função do pacote rand para aleatorizar um numero de 0 a 99 mas com a soma de 1 para que fique 1 a 100
	target := rand.Intn(100) + 1

	var guess int
	var tentativas int

	for {
		fmt.Print("Digite um numero de 1 a 100 ou 0 para sair: ")
		fmt.Scan(&guess)

		if guess == 0{
			fmt.Println("Você desistiu, o numero era:", target)
		} else if guess < target {
			fmt.Println("Um pouco mais")
			tentativas++
		} else if guess > target {
			fmt.Println("Um pouco menos")
			tentativas++
		} else if guess == target {
			fmt.Printf("Parabens voce acertou! com %d tentativas \n", tentativas)
			return
		}

	}

}