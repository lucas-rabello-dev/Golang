# Generics em Golang

resumo de um vídeo do canal <a href="https://www.youtube.com/@huncoding">HunCoding</a> 
<a href="https://www.youtube.com/watch?v=nD_6L2PBfaE">clique aqui para ir ao vídeo</a>

<hr>

## O que são Generics?
A programação genérica é um estilo de programação no qual os algoritimos são escritos em tipos a serem especificados posteriormente.

Sintaxe de uma função genérica:
```
func <nome_da_função>[<nome_tipo> <restrição_tipo>](<parâmetro> <tipo>, ...) tipo_retorno {
    <código_da_função>
}
```
Código em Go:

```go
func Max[T int | float64](a, b T) T {
	if a > b {
		return a
	}
	return b
}
```
Ao lado do nome da função, o que vem entre `[ ]` no caso `[T int | float64]` é a declaração do parâmetro de tipo, ou seja, aqui você define os tipos que a função vai aceitar.

`T` -> É o nome da váriavel de tipo
`int | float64` -> Go só permite que `T` seja `int` ou `float64`.

`a, b T` são os parâmetros normais da função que são do tipo `int` ou `float64` nesse caso.

### Como chamar a função

```go
Max(20, 10) // T = int
Max(3.5, 6.5) // T = float64
```
Você pode também forçar o tipo:
```go
Max[int](20, 10)
Max[float64](3.5, 6.5)
```

<span style="color: red;"><i><b>Os tipos dos parâmetros tem que ser iguais!</b></i></span>




