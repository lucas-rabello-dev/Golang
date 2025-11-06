package main

import "fmt"

func main() {
	var nums = [3]int{0, 1, 2}
	var ptr *int = &nums[0]
	fmt.Println(ptr, *ptr)
	Sv()
	
}

// segmentation violation
func Sv() {
	var ptr *int = nil
	a := *ptr
	if a == 0 {}
}