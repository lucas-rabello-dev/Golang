package main

import (
	"math/rand"
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	for i := 0; i < 5; i++ {
		num := rand.Intn(5) + 1
		fmt.Printf("Rotina: %d Num: %d\n", i, num)
	}

	for i := 0; i < 5; i++ {
		num := rand.Intn(5) + 1
		fmt.Printf("Rotina: %d Num: %d\n", i, num)
	}

	fmt.Println("main acabou")
	fmt.Println(time.Since(now))
	
}