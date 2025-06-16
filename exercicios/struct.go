package main


import "fmt"


type Cadastro struct {
	Nome string
	Idade uint
	Salario map[bool]uint // se sim e o salario
	Emprego map[bool]string // se sim e o nome
}


func main() {

	fmt.Println("---- CADASTRO ----")

	var (
		nome string
		idade uint
		trabalho string
		veri_salario bool
		veri_emprego bool
		salario uint
		emprego string
	)

	salario_map := make(map[bool]uint)
	emprego_map := make(map[bool]string)

	fmt.Print("Digite o nome: ")
	fmt.Scan(&nome)

	fmt.Print("Digite a idade: ")
	fmt.Scan(&idade)

	fmt.Print("O Cliente trabalha (s) (n): ")
	fmt.Scan(&trabalho)

	

	if trabalho == "s" {
		fmt.Scanln() // para limpar o buffer (pra variar kkkkkkkkkk)
		fmt.Print("Qual o nome do emprego: ")
		fmt.Scanln(&emprego)

		fmt.Scanln() // para limpar o buffer (pra variar kkkkkkkkkk)
		fmt.Print("Qual o salário do cliente: ")
		fmt.Scan(&salario)

		veri_emprego = true
		veri_salario = true
	} else if trabalho == "n" {
		veri_emprego = false
		veri_salario = false
		emprego = "Desempregado"
		salario = 0
	} else {
		fmt.Println("Resposta inválida!")
	}

	salario_map[veri_salario] = salario
	emprego_map[veri_emprego] = emprego



	cadastro := Cadastro {
		Nome: nome,
		Idade: idade,
		Salario: salario_map,
		Emprego: emprego_map,
	}
	

	fmt.Println("Cadastro finalizado: ", cadastro)



}