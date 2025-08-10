package main

import (
	"contas/utils"
	"fmt"
	"sync"
)
const rangeNum int = 100

func main() {

	chSend := make(chan int)
	chRead := make(chan int)

	var wg sync.WaitGroup
	wg.Add(3)

	num := utils.GenerateNum(rangeNum)

	// CASO NAO TIVESSE ESSA FUNC TERIA ERRO DEADLOCK
	// Goroutine "broker" que lê do chSend e repassa para chRead
	go func() {
		for val := range chSend {
			chRead <- val
		}
		close(chRead)
	}()

	go func() {
		defer wg.Done()
		var plus int = utils.Plus(rangeNum, num)
		chSend <- plus // enviando
		var plusValue = <- chRead // recebendo
		fmt.Printf("número gerado inicialmente: %d | número da soma gerado: %d \n", num, plusValue)
	}()

	go func() {
		defer wg.Done()
		var minus int = utils.Minus(rangeNum, num)
		chSend <- minus
		var minusValue = <- chRead
		fmt.Printf("número gerado inicialmente: %d | número da subtração gerado: %d \n", num, minusValue)
	}()

	go func() {
		defer wg.Done()
		var times int = utils.Times(rangeNum, num)
		chSend <- times
		var timesValue = <- chRead
		fmt.Printf("número gerado inicialmente: %d | número da multiplicação gerado: %d \n", num, timesValue)
	}()

	wg.Wait()
	close(chSend)
}