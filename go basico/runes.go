package main

import (
	"fmt"
)

func main() {
	teste := `ss` // string raw literal
	teste2 := 's' // char qunado usa aspas simples tbm pode ser chamado de rune
	var r rune = 'a'
	var (
		t rune = 't'
		w rune = 'w'
		b rune = 'b'
	)
	fmt.Println(r, t, w, b) // 97 (valor numérico Unicode)
	fmt.Printf("%c\n", r)   // a
	// %c -> imprime como um char ASCII
	fmt.Println("%T %T",teste2, teste)
}