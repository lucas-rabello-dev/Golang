package main

import ( 
	"fmt"
	"strings"
)

func main() {

	text := "go é uma linguagem de programação e go é muito rápida"

	words := strings.Fields(text)

	wordCount := make(map[string]int)

	for _, word := range words {
		wordCount[word]++ // adiciona no map o a palavra e a quantidade que ela repete
	}

	fmt.Println("Contagem de palavras: ")

	for word, count := range wordCount {
		fmt.Printf("Palavra: %s Frequência: %d \n", word, count)
	}

}