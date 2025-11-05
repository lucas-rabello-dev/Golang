package main


import (
	"fmt"
	"time"
)


func main() {

	// inicia o tempo atual
	currentTime := time.Now()
	// pega a hora
	hour := currentTime.Hour()
	// pega o dia
	day := currentTime.Weekday()


	fmt.Println(currentTime, "(---)", hour, day)
}