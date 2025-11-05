package ip

import (
	"fmt"
	"testing"
)

func TestIp(t *testing.T) {
	var result bool = Ip(22)
	if result == true {
		fmt.Println("Esse número é par!")
		return
	}
	t.Error("Esse número é impar!")
}