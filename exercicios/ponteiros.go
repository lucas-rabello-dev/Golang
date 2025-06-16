package main


import "fmt"

func main() {


	var num int = 10

	var ptr *int = &num
	// ou assim
	ptr2 := new(int)


	fmt.Println(ptr)


}