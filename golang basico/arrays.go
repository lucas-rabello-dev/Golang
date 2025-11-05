package main 



import (
	"fmt"
)


func main() {

	var notas = [5]float64{3.3, 6.7, 7.4, 9.9, 8.9}


	soma := notas[0] + notas[1] + notas[2] + notas[3] + notas[4]

	media := soma / float64(len(notas))

	fmt.Println("A médias das notas é:", media)

}