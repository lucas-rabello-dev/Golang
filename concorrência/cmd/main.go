package main

import (
	"buscador/internal/fetcher"
	"fmt"
	"sync"
	"time"
)


func main() {
	start := time.Now()

	priceChannel := make(chan float64)

	go func() {
		var totalPrice float64
		countPrice := 0.0
		for price := range priceChannel {
			totalPrice += price
			countPrice++
			fmt.Printf("R$ %.2f \n", price)
			fmt.Printf("Preço Médio: %.2f \n", totalPrice/countPrice)
		}
	}()

	var wg sync.WaitGroup // grupo de espera (wait group)
	wg.Add(3) // 3 goroutines 

	// funções anônimas
	go func() {
		defer wg.Done() // executa quando a função for executada
		priceChannel <- fetcher.FetchPriceFromSite1()
	}()

	go func() {
		defer wg.Done() // executa quando a função for executada
		priceChannel <- fetcher.FetchPriceFromSite2()
	}()

	// colocar o go na frente diz para rodar em paralelo
	go func() {
		defer wg.Done() // executa quando a função for executada
		priceChannel <- fetcher.FetchPriceFromSite3()
	}() // -> significa que a função vai ser executada assim que for criada

	wg.Wait() // aguarda todas as goroutines terminarem
	close(priceChannel)

	// tempo de excução do códgo
	fmt.Printf("\n Tempo total: %s", time.Since(start))
	
}