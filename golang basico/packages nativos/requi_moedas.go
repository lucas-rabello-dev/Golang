package main


import (
	"fmt"
	"io"
	"net/http"
)


func main() {
	URL:= "https://economia.awesomeapi.com.br/last/USD-BRL,EUR-BRL,BTC-BRL"

	resp, err := http.Get(URL)

	if err != nil {
		fmt.Println("Houve um erro:", err)
		return
	}

	defer resp.Body.Close()


	body, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("Houve um erro:", err)
		return
	}

	fmt.Println("Conteudo da API:", string(body))
}