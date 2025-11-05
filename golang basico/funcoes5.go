package main


import "fmt"


// função que recebe um ponteiro com argumento e altera o valor da variavel 
// original

func alterValor(x *int) {
	*x *= 2
}


func main() {
	num := 5

	fmt.Printf("Valor inicial: %d \n", num)

	alterValor(&num)

	fmt.Printf("Valor após a alteração: %d \n", num )
	
}