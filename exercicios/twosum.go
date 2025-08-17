package main

import "fmt"

func twoSum(array []int, target int) []int {

	var final_array []int

	for i := 0; i < len(array); i++ {
		for j := i + 1; j < len(array); j++ {
			if array[i] + array[j] == target{
				final_array = append(final_array, i, j)
				return final_array
			}
		}
	}
	final_array = append(final_array, -1, -1)
	return final_array	
}
func main() {

	array := []int{2, 4, 6, 2}

	target := 4

	resul := twoSum(array, target)
	fmt.Printf("Valores: %d + %d \n", array[resul[0]], array[resul[1]])
	fmt.Printf("Indices: %d e %d é igual a: %d \n", resul[0], resul[1], target)
}