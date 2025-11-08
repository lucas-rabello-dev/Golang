package main

/*
#cgo LDFLAGS: -L. -llib
extern void HelloFromCPP();
*/
import "C"

func main() {
    C.HelloFromCPP()
}
