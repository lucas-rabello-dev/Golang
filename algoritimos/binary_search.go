package main

import (
	"fmt"
	"time"
)

func BinarySearch(array []int, target int) (int, int, int) {
	tentativas := 0
	firstValue := 0
	lastValue := len(array) -1

	for i := 0; i < len(array); i++{
		halfValue := (firstValue + lastValue) / 2
		kick := array[halfValue]
		if kick == target {
			return array[halfValue], tentativas, halfValue // retorna o valor, tentativas e o indice
		} else if kick > target {
			lastValue = halfValue - 1
		} else {
			firstValue = halfValue + 1
		}
		tentativas += 1
	}
	return -1, tentativas, -1
}

func main() {
	now := time.Now()
	fmt.Println("O código foi iniciado!")

	array := []int{}
	for i := 1; i <= 1000; i++ {
		array = append(array, i)
	}
	fmt.Println(array)
	value, ten, index := BinarySearch(array, 2)
	if value == -1 {
		fmt.Printf("O valor: %d não foi encontrado! \n", value)
		return
	}
	fmt.Printf("O valor: %d | Indice: %d foi encontrado com %d tentativas\n", value, index, ten)
	fmt.Println("tempo de execução:", time.Since(now))
}