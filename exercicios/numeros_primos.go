// uma função que retorne apartir de um input os numeros primos
// apenas divisivel por 1 e por ele mesmo (numeros primos)

package main


import "fmt"


func NumPrimo(num int) bool {
	if num < 2 {
		return false
	}

	for i := 2; i*i <= num; i++ {
		if num % i == 0 {
			return false
		}
	}
	return true
}

func main() {


	var num int


	fmt.Print("Digite o numero para ver se ele é primo: ")
	fmt.Scan(&num)

	var funcao bool = NumPrimo(num)

	if funcao == true {
		fmt.Printf("O numero %d é um numero primo! \n", num)
	} else if funcao == false {
		fmt.Printf("O numero %d não é um numero primo! \n", num)
	}



}