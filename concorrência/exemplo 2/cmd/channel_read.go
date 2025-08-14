package main


import (
	"sync"
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)
	var wg sync.WaitGroup

	// Criando 3 goroutines que vão ler o canal
	for i := 0; i <= 6; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			for msg := range ch {
				fmt.Printf("Goroutine %d recebeu %s \n", id, msg)
				time.Sleep(time.Second * 500)
			}
		}(i)
	}
	// enviando msg para o canal
	messages := []string{"Olá", "mundo", "go", "concorrência", "é", "incrível"}
	for _, msg := range messages {
		ch <- msg
		time.Sleep(time.Millisecond * 200)
	} 

	close(ch)

	wg.Wait()
	fmt.Println("cabou")


}