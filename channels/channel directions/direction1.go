package main

import "fmt"

// como o "chan<-" ele só pode receber valores
func ping(pings chan<- string, msg string) {
	pings <- msg
}
// "<-chan" pode enviar valores para o channel
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings // pings(um channel) esta enviando seus valores para "msg"
	pongs <- msg // pongs(um channel) está recebendo
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message") // pega a msg e coloca no channel
	pong(pings, pongs) // pings(envia dados) e pongs(recebe dados)
	fmt.Println(<-pongs) // Envia dados para o Println()
}