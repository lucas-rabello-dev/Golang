package main

import (
	"fmt"
	"math/rand"
)

// ------------------------------------------- Exemplo de Interfaces -----------------------------------------------
// Interfaces é um tipo que definem um conjunto de métodos que um tipo pode ter

// Explicação do Chatgpt

// Pensa nela como um "contrato" que diz:
// “Se você tiver estes métodos, eu te considero um desse tipo.”


// Interface 
type Animal interface {
	Latir() string // -> definindo que animal terá um método latir e que retorne uma string
}

// Tipos que implementam a interface
type Cachorro struct{
	Nome string "Nome do cachorro"
}

// essa função faz parte da struct Cachorro e seu método
func (Cachorro) Falar() string { 
	return "Au au!" 
}

// Função que aceita qualquer tipo que implemente Animal
func FazerAnimalFalar(a Animal) {
	fmt.Println(a.Latir())
}


// ------------------------------------------- Exemplo de methods ---------------------------------------------------
type Pessoa struct {
	Nome string
	Idade int
	CPF string
}
// reciver de valor -> não altera o valor original
func (p Pessoa) Info() {
	fmt.Printf("Nome: %s | Idade: %d | CPF: %s \n", p.Nome, p.Idade, p.CPF)
}

func (p *Pessoa) UpdateAge() {
	p.Idade = rand.Intn(60)
	fmt.Printf("Nova Idade: %d \n", p.Idade)
}

func main() {

	// ------------------------------------------- Exemplo de methods ---------------------------------------------------
	// Adicionando apenas uma pessoa por vez
	// p := Pessoa{
	// 	Nome: "lucas",
	// 	Idade: 16,
	// 	CPF: "00000000000",
	// }
	
	// p.Info()
	// p.UpdateAge()
	// p.Info()

	// Criando multiplos usuários e listando cada um com for
	// p := []Pessoa{
	// 	{Nome: "lucas", Idade: 21, CPF: "0982485732"},
	// 	{Nome: "rafael", Idade: 44, CPF: "777837467"},
	// }

	// for _, pessoa := range p {
	// 	pessoa.Info()
	// } 

	// ------------------------------------------- Exemplo de Interfaces -----------------------------------------------


	c := Cachorro{}
	fmt.Println(c.Falar())



}