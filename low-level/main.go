package main

import (
	"fmt"
	"unsafe"
)

func main() {
	arr := [3]byte{'a', 'b', 'c'}

	fmt.Println(unsafe.Sizeof(arr), "bytes") // 3 bytes
	fmt.Println(int(unsafe.Sizeof([]byte("sla"))) / 8, "bytes") // retorna 3 bytes
	fmt.Println(unsafe.Sizeof(byte('a')), "byte") // 1 bytes
}

// segmentation violation
func Sv() {
	var ptr *int = nil
	a := *ptr
	if a == 0 {}
}