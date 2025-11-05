package main


import (
	"fmt"
	"strconv"
)



func main() {

	// convertendo str para int
	numFake := "1234"
	numero, _ := strconv.Atoi(numFake) // essa função tem erro

	fmt.Println(numero)

	// int para str
	numeros := 2312
	numeros_str := strconv.Itoa(numeros)

	fmt.Println(numeros_str)

	// str para float
	floatStr := "12.44"
	floatNormal, _ := strconv.ParseFloat(floatStr, 64)

	fmt.Println(floatNormal)

}