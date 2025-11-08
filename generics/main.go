// simple example of generics
package main

import (
	"fmt"
)
func Soma[T int | float64](a, b T) T {
	return a + b
}

func main() {

	fmt.Println(Soma(5, 10.50))
}