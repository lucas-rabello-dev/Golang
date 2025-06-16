package main


import (
	"fmt"
	"net/http"
	"io"
)

func main() {
	URl := "https://jsonplaceholder.typicode.com/todos/2"

	// duas variáveis que são uma a resposta e a outra o erro
	resp, err := http.Get(URl)

	if err != nil {
		fmt.Println("Erro ao fazer a requisição!", err)
		return
	}
	// quando o código acabar ele vai finalizar a resposta
	defer resp.Body.Close()

	// Usando io.ReadALL para ler o corpo da resposta

	// o corpo e o erro
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Erro ao ler o corpo da resposta: ", err)
		return
	}
	// converteu a resposta para uma string
	fmt.Println("Resposta da API:", string(body))

}