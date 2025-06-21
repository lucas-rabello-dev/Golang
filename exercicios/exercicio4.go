package main

import (
	"fmt"
)

func romanToInt(s string) int {
    for _, i := range s {
        fmt.Println(i)
    }
}

func main() {
	romanToInt("ddea")
}

