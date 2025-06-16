package main


import (
	"fmt"
	"os"
)


func main() {

	// criando um arquivo
	arquivo, err := os.Create("exemplo.txt")

	if err != nil {
		fmt.Println("Houve um erro ao criar o arquivo:", err)
		return
	}

	// para fechar o arquivo no final da execução
	defer arquivo.Close()

	_, err = arquivo.WriteString("Texto de exemplo no arquivo.")

	if err != nil {
		fmt.Println("Houve um erro ao escrever no arquivo:", err)
		return
	}

	fmt.Println("Arquivo criado e texto escrito com sucesso!")


	// lendo um arquivo

	conteudo, err := os.ReadFile("exemplo.txt")

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Conteudo:", string(conteudo))



}