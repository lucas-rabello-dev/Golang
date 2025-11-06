package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	now := time.Now()
	var wg sync.WaitGroup
	
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			num := rand.Intn(5) + 1
			fmt.Printf("Goroutine: %d Seu valor: %d \n", id, num)
			defer wg.Done()
		}(i)
		// time.Sleep(time.Second * 1)
	}
	

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			num := rand.Intn(5) + 1
			fmt.Printf("Goroutine: %d Seu valor: %d \n", id, num)
			defer wg.Done()
		}(i)
		// time.Sleep(time.Second * 1)
	}

	wg.Wait()
	fmt.Println("main acabou", time.Since(now))

}