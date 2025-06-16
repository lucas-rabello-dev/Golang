package main


import "fmt"

type Cliente struct{
	Nome string
	Idade int
}

// função que recebe um ponteiro para alterar os dados do cliente 

func atualizarCliente(c *Cliente, novoNome string, novaIdade int){
	c.Nome = novoNome
	c.Idade = novaIdade
}


func main() {
	cliente := Cliente{Nome: "Lucas", Idade: 22}
	fmt.Printf("Antes da alteração: Nome: %s Idade: %d \n", cliente.Nome, cliente.Idade)
	atualizarCliente(&cliente, "Carlos", 32)
	fmt.Printf("Depois da alteração: Nome: %s Idade: %d \n", cliente.Nome, cliente.Idade)



}