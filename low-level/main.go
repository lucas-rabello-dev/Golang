package main

import "fmt"

type Sim struct {
	item int
}

func main() {
	p := new(Sim)
	fmt.Println(p, *p)

}