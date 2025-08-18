# Concurrency in Golang
resumo do vídeo do <a href="https://www.youtube.com/@huncoding">HunCoding</a>

<a href="https://www.youtube.com/watch?v=PMM5Fia9O4g">clique aqui para ir ao vídeo</a>

<hr>

## O que é um processo?
Processo é uma instância de um programa que está rodando

<hr>

### O que tem dentro de um processo?
- Sistema Operacional aloca memória
- <b><i>Code</b></i> - Instruções de máquina
- <b><i>Data</b></i> - Valores/Variaveis globais
- <b><i>Heap</b></i> - Alocação de memória dinâmica
- <b><i>Stack</b></i> Utilizado para guardar variáveis locais de uma função

<hr>

## O que são Threads?
<big>
<p><b><i>Threads</i></b> são a <b><i>menor unidade de execução</b></i> que a CPU aceita</p></big>

- Cada processo tem no mínimo uma thread (que é a thread principal)
- Um processo pode ter vários threads
- Threads compartilham o mesmo espaço de endereço
- Threads executam uma independente da outra
- Threads podem rodar em concorrência ou paralelamente 

<hr>

### Estados do thread

<table>
<tr>
    <th>RUNNEBLE</th>
    <th></th>
    <th>SCHEDULER</th>
    <th></th>
    <th>RUNNING</th>
    <th></th>
    <th>WAITING</th>
</tr>
<tr>
    <td>Quando um processo é criado, a thread principal é colocada na fila como pronto para executar.</td>
    <td>➡️</td>
    <td>Assim que a CPU fica disponível ele leva a thread para o terceiro estado.</td>
    <td>➡️</td>
    <td>Assim que a CPU fica disponível, a thread começa a executar em um espaço de tempo.</td>
    <td>OU</td>
    <td>Se a thread ficar bloqueada por conta de operações I/O (input e output) como ler/escrever em disco, conexões ou esperando eventos de outros processos.</td>
<tr>
</table>
<p><big>Se o tempo definido espirar, a thread é colocada de novo na fila</big></p>

<hr>

## Limitações de dividir processos e threads para a concorência em Go

### Context switching
A CPU gasta tempo copiando o contexto da thread atual em execução em memória e resgatando o contexto das próximas threads à executar. A CPU não está fazendo context switching (troca de contexto) 

Portanto, é mais eficiênte usar um processo que contém várias threads.

### C10K Problem

O scheduler do sistema aloca um espaço de tempo para a execução de cada processo

Esse tempo definido é dividido igualmente entre cada thread executada dentro desse processo.

<table>

<tr> 
    <th>Scheduler Period</th>
    <th>Number of threads</th>
    <th>Thread time slice</th>
</tr>

<tr>
    <td>10ms</td>
    <td>2</td>
    <td>5ms</td>
</tr>

<tr> 
    <td>10ms</td>
    <td>5</td>
    <td>2ms</td>
</td>

<tr> 
    <td>10ms</td>
    <td>5</td>
    <td>10us?</td>
</td>

</table>

Quanto mais thread, menos tempo cada uma vai ter para executar e assim, possivelmente a troca de contexto do sistema vai demorar mais para acontecer do que propriamente o tempo que cada thread vei ter para executar.

## Mas onde entram as goroutines?