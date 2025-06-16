package main

import (
	"fmt"
)


func indice(lista []string, add string) int {
	
	for indice, valor := range lista {
		if add == valor {
			return indice
		} else {
			fmt.Printf("A tarefa %s não foi adicionada a lista ainda! \n", add)
		}
	}
	return 0 // retorna 0 para erro 
}


func add_item(lista []string, item string) []string{

	lista = append(lista, item)
	return lista

}


func remove_item(lista []string, item string) []string {

	indice_ := indice(lista, item)

	if indice_ == 0 {
		fmt.Println("Houve um erro! \nNão foi possivel encontrar:", item, "na lista")
		return lista
	}

	lista = append(lista[:indice_], lista[indice_ + 1:]...)
	return lista 
}

func ver_itens(lista []string) {

	if len(lista) == 0 {
		fmt.Println("Nada adicionado ainda!")
		return
	}

	for _, i := range lista {
		fmt.Println(i)
	}
}

func main() {
	lista := []string{}


	var remove string
	var add string

	var input int

	for {
		fmt.Print("Comandos: (1) Ver a lista (2) Adicionar tarefa (3) Remover tarefa (4) Sair: ")
		fmt.Scan(&input)
		if input == 4 {
			fmt.Println("Saindo...")
			return
		} else if input == 1 {
			ver_itens(lista)
		} else if input == 2 {
			fmt.Print("Digite o que você quer adicionar na lista: ")
			fmt.Scan(&add)
			nova := add_item(lista, add)
			lista = nova
		} else if input == 3 {
			fmt.Print("Digite o que você quer remover: ")
			fmt.Scan(&remove)
			nova := remove_item(lista, remove)
			lista = nova
		}
	}
}


// métodos
	// append  lista = append(item)

	// remover
	// lista = append(lista[:indice_item_remove], lista[indice_item_remove + 1:])