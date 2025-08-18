// Você deve criar um programa que leia vários arquivos de texto e conte
// quantas vezes cada palavra aparece, usando concorrência para acelerar o processamento.

package main

import (
	"os"
	"strings"
	"sync"
	"fmt"
	"time"
)
const filePath string = "./textos/" // para evitar conflito na hora de abrir os arquivos 

func main() {
	inicio := time.Now()

	// carrega o diretório no início
	files, err := os.ReadDir("./textos")
	if err != nil {
		panic(err)
	}
	var wg sync.WaitGroup
	wg.Add(len(files)) // espera o numeros de files para terminar
	ch := make(chan map[string]int, len(files)) // Adicionando um buffer 
	// O buffer do canal serve para que as goroutines possam 
	// enviar valores mesmo que ninguém esteja recebendo naquele momento.
	// Ou seja, pode guardar até 3 valores diferentes que no caso são 3 maps
	
	for _, file := range files {
		name := file.Name()
		go func(name string) {
			defer wg.Done()


			openFile, err := os.Open(filePath + name) // Usado o Open para apenas a leitura
			if err != nil {
				panic(err)
			}

			content, err := os.ReadFile(filePath + name)
			if err != nil {
				panic(err)
			}

			words := strings.Fields(string(content))
			counts := make(map[string]int) // está aqui pois cada goroutine deve ter o seu proprio map para evitar erro
			for _, word := range words {
				counts[word]++ // No caso aqui ele cria dentro do map a palavra e incrementa 1 a cada vez que ela aparecer
			}
			ch <- counts // aqui é enviado todas as palavras juntas ao invés de um par de cada vez

		defer openFile.Close() // fechar o arquivo ao final da goroutine
		}(name)
	}
	wg.Wait()
	finalWords := make(map[string]int)
	for i := 0; i < len(files); i++ {
		m := <- ch
		for word, count := range m {
			finalWords[word] += count
		}
	}
	for key, value := range finalWords {
		fmt.Printf("Palavra: %s | Número de vezes que aparece nos arquivos: %d \n", key, value)
	}
	fmt.Println(time.Since(inicio)) // mostra o tempo de execução
}