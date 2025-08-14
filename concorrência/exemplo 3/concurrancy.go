package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	var wg sync.WaitGroup

	nums := make([]int, 10)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			for j := 0; j < 1000000000; j++ {
				nums[i]++
			}
			fmt.Printf("num%d: %d\n", i+1, nums[i])
		}(i)
	}

	wg.Wait()
	fmt.Printf("Tempo total: %s\n", time.Since(start))
}
