package main

import (
	"fmt"
)



func main() {
	i := 23

	switch i {
	case 25:
		fmt.Println(25)
		fmt.Println("Fim do code 2")
	case 23: 
		fmt.Println("Vinte e três")
		fallthrough
	default: 
		fmt.Println(nil)
	} 

	fmt.Println("fim do code 1")
}