package main

import (
    "fmt"
)

func main() {

    // ASCCI/UTF-8 values
    data := []byte("hello") // convert string to raw bytes
    fmt.Println(data)

    s := "lucas"
    b := []byte(s) // convert string to bytes
    fmt.Println(b)
    s2 := string(b) // convert bytes back to string
    fmt.Println(s2)

    // Why use bytes
    // You can use to modify strings
    b2 := []byte("hello")
    fmt.Printf("Before: %s\n", string(b2))
    b2[0] = 'H'
    fmt.Printf("Modified: %s\n", string(b2))

}