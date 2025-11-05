package main

import "fmt"


// func 1
func welcome() {
	fmt.Println("Bem vindo ao sistema de filmes")
}

// funnc 2


func createMovie() {
	var name string
	var yearRelease int
	var moviePrice float64

	fmt.Print("Digite o nome do filme: ")
	fmt.Scanln(&name)

	fmt.Print("Digite o ano de lançamento: ")
	fmt.Scanln(&yearRelease)

	fmt.Print("Digite o preço do filme: ")
	fmt.Scanln(&moviePrice)

	fmt.Printf("%s (%d) - R$ %.2f \n", name, yearRelease, moviePrice)


}

func main() {
	fmt.Println("Func principal")
	createMovie()
	
}
