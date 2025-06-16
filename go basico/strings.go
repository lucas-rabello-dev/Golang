package main



import (
	"fmt"
	"strings"
)


func main() {

	


	texto_exemplo := "coisas bem legais, eu programo em go Mano"
	new := strings.ToLower(texto_exemplo)

	if strings.Contains(new, "mano") {
		fmt.Println("Tem")
	} else {
		fmt.Println("Não tem")
	}

	




}
