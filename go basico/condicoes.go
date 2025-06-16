package main


import (
	"fmt"
)


func main() {

	
	var idade int

	fmt.Print(">>> ")
	fmt.Scan(&idade)


	if idade >= 18 {
		fmt.Println("maior de idade")
	} else if idade < 18 {
		fmt.Println("menor de idade")
	}


	



}


