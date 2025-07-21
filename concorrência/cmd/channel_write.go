package main

import (
	"time"
	"fmt"
)

func main() {
	ch := make(chan int)
	
	// Iniciando uma goroutine que envia números para o canal

	go func() {
		for i := 0; i <= 5; i++{
			fmt.Printf("Enviando número %d para o canal \n", i)
			ch <- i
			time.Sleep(time.Second)

		}
		close(ch)
	}()

	for num := range ch {
		fmt.Printf("Número recebido %d \n", num)
	}
	
}