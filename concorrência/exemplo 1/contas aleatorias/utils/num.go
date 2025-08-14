package utils


import (
	"math/rand"
)

// aleatoriza um numero com base em um range
func GenerateNum(range_num int) int{
	return rand.Intn(range_num)
}


func Minus(range_num, num int) int {
	return rand.Intn(range_num) - num
}

func Plus(range_num, num int) int {
	return rand.Intn(range_num) + num
}

func Times(range_num, num int) int {
	if range_num % 2 == 0{
		second_num := range_num / 2
		return rand.Intn(range_num) * second_num
	}
	second_num := range_num / 3 
	return rand.Intn(range_num) * second_num
}