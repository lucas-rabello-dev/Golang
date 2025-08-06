package main

import (
    "fmt"
)

func main() {
    var letras string = "abcdefghijklmnopqrstuvwxyz"
    // var lertasM string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
    for index, r := range letras {
        fmt.Printf("Index: %d | Char: %c | Rune (dec): %3d |Rune: %U \n", index, r, r, r)
    }
}
