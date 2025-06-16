package main 

import "fmt"



type Cadastro struct {
	Nome string
	Telefone string
	Email string
}


func main() {

	fmt.Println("Sistema de cadastro em golang!")

	pessoa1 := Cadastro{
		Nome: "lucas",
		Telefone: "11 99999-9999",
		Email: "lucas@teste",
	}

	dados := map[string]string{
		"Nome":pessoa1.Nome ,
		"Telefone":pessoa1.Telefone,
		"Email":pessoa1.Email,
	}

	for chave, valor := range dados {
		fmt.Printf("%s: %s \n", chave, valor)
	}

}