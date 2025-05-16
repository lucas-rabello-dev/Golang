package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	first, err1 := strconv.Atoi(os.Args[1])

	if err1 != nil{
		fmt.Println("Error:", err1)
		return
	}

	second, err2 := strconv.Atoi(os.Args[2])

	if err2 != nil{
		fmt.Println("Error:", err2)
		return
	}

	third, err3 := strconv.Atoi(os.Args[3])

	if err3 != nil{
		fmt.Println("Error:", err3)
		return
	}

	var product int = first * second

	if product == third{
		fmt.Println("The product of", first, "*", second, "=", third)
	} else if product > third{
		fmt.Println("The product of", first, "*", second, "is greater than", third)
	} else if product < third{
		fmt.Println("The product of", first, "*", second, "is less than", third)
	}

}