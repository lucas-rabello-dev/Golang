### como Go chama funções de C++ usando CGO
O Go chama código compilado em C++ através de uma biblioteca estática (.a), e o cgo é o tradutor entre eles.

o comando do C++ no makefile:
`g++ -c -o lib.o main.cpp`
ele compila mas o `-c` faz com que não seja linkado
e isso gera um arquivo objeto (`.o`) 
`-o lib.o` define o nome de arquivo de saida

Criando a biblioteca estática
`ar rcs liblib.a lib.o`

`ar` → archiver, empacota objetos (.o) em uma biblioteca .a.

`r` → substitui arquivos se já existirem dentro do pacote.

`c` → cria o arquivo se não existir.

`s` → cria um índice (pra linkagem mais rápida).

`liblib.a` → é a biblioteca estática final.

`lib.o` → o conteúdo dela.


```cpp
#include <iostream>
extern "C" {
    void HelloFromCPP() {
        std::cout << "Olá do C++!" << std::endl;
    }
}
```
`#include <iostream>` → inclui as funções de I/O padrão (como std::cout).

`extern "C"` → desativa o name mangling do C++.
(Sem isso, o nome real da função viraria algo feio tipo _Z13HelloFromCPPv, que o Go não reconheceria.)

`void HelloFromCPP()` → função comum em C, exportada pra ser chamada pelo Go.
 