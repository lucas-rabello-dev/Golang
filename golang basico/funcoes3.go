package main

import (
	"fmt"
)


func teste(num int) {
	if num >= 100{
		return
	}

	num++
	fmt.Println(num)
	teste(num)
}


func main() {
	
	num := 0

	teste(num)
}

