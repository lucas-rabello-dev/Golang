package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

const filePath string = "./textos/" // para evitar conflito na hora de abrir os arquivos

func main() {
	inicio := time.Now()

	// carrega o diretório no início
	files, err := os.ReadDir(filePath)
	if err != nil {
		panic(err)
	}

	finalWords := make(map[string]int)

	// percorre cada arquivo de forma sequencial
	for _, file := range files {
		name := file.Name()

		openFile, err := os.Open(filePath + name)
		if err != nil {
			panic(err)
		}
		defer openFile.Close()

		content, err := os.ReadFile(filePath + name)
		if err != nil {
			panic(err)
		}

		// separa as palavras
		words := strings.Fields(string(content))

		// conta quantas vezes cada palavra aparece
		for _, word := range words {
			finalWords[word]++
		}
	}

	// exibe os resultados
	for key, value := range finalWords {
		fmt.Printf("Palavra: %s | Número de vezes que aparece nos arquivos: %d\n", key, value)
	}

	fmt.Println("Tempo de execução:", time.Since(inicio))
}

