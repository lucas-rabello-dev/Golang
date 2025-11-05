package main


import "fmt"


func main() {

	// o primeiro item (2) são as colunas e o segundo (3) são as linhas
	a := [2][3]int{
		{0, 1, 2},
		{3, 4, 5},
	}

	fmt.Println(a)

}