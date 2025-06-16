package main

import "fmt"

// função para soma de números com váriadicos
func Sum(nums ...int) {
	sumTotal := 0
	for _, i := range nums {
		sumTotal += i
	}
	fmt.Printf("Soma é: %d \n", sumTotal)
}


// função para apresentação de cursos com variádicas
func apresentação(data map[string]string) {
	for key, value := range data {
		fmt.Printf("%s - %s \n", key, value)
	}
}

func main() {

	apresentação(map[string]string{
		"name": "python",
		"category": "back-end",
		"level": "iniciante",
	})

	apresentação(map[string]string{
		"name": "python para gênios",
		"category": "AI",
		"level": "Avançado",
	})

	apresentação(map[string]string {
		"name": "python CLI",
		"category": "back-end 2.0",
		"level": "master",
	})





}