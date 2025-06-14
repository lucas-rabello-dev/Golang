package main



import "fmt"



func main() {

	// map com var e tem inicializar colocando os valores ou usando o make()
	var dados map[string]int  // apenas declarado (nil)
	dados = make(map[string]int) // agora pode usar

	// inicialização com make()
	dados2 := make(map[string]int)

	fmt.Println(dados2)

	// criação com valor 

	dados3 := map[string]int{
		"teste1": 10,
		"teste2": 20,
	}

	// acessando e modificando valores

	dados["TESTE"] = 900 // adiciona ou atualiza

	idade := dados3["teste1"] // no caso acessa o valor 10 que é atribuido a variável idade
	fmt.Println(idade)


	delete(dados3, "teste2") // deleta a chave e o valor

	for nome, num := range dados3 {
		fmt.Printf("Nome: %s Num: %d \n", nome, num)
	}


	// ver se um valor existe dentro de um map

	mapTeste := make(map[string]int)

	mapTeste["TESTE NO MAP"] = 999
	mapTeste["123"] = 888

	Chave, existe := mapTeste["123"]

	// se Chave que é o "123" existir ele retorna o seu valor 
	// se não para int retorna 0 e para str retorna ""

	// existe verifica se existe e é um bool

	if existe {
		fmt.Println("Existe:", Chave) // retorna o valor da chave não a chave
	} else {
		fmt.Println("Não existe")
	}
 



}