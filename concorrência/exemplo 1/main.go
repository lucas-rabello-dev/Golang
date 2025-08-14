package main


import (
	"fmt"
	"sync"
)

// Os channel precisam receber e enviar dados nas goroutines para evitar o DeadLock -> erro
// e wg.Wait precisa estar depois das goroutines para evitar esse mesmo erro

func main() {
	// criando um channel de strings com make
	ch := make(chan string)

	var wg sync.WaitGroup // criando o grupo de espera
	wg.Add(2) // adiciona 2 goroutines



	var teste string = "nome de teste"
	// go -> indica que a func (anônima) vai rodar em paralelo com o código
	go func() {
		defer wg.Done()
		ch <- teste
		fmt.Println(teste) // Printa o valor da váriavel
	}() // -> significa que a função vai ser executada assim que for criada
	go func() {
		defer wg.Done()
		var segundo_ch string = <- ch
		fmt.Println(segundo_ch) // Printa o valor da váriavel
	}() // -> significa que a função vai ser executada assim que for criada


	wg.Wait() // aguarda todas as goroutines terminarem

	close(ch) // fecha o channel no final do código
}