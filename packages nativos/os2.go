package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)




// 1- retornar a pasta atual
func getCurrentDirectory() string {
	dir, err := os.Getwd()

	if err != nil {
		fmt.Println(err)
		return "NONE"
	} else {
		return dir
	}
}

// 2- Listar aquivos e pastas
func listFilesAndDicts() {
	files, err := os.ReadDir(".") // uso do ponto para ler todos os arquivos

	if err != nil {
		fmt.Println(err)
		return
	}
	for _, file := range files {
		fmt.Println(file.Name())
	}
}

// 3- ver versão do SO
func goOSVersion() {
	var cmd *exec.Cmd

	if runtime.GOOS == "windows"{
		cmd = exec.Command("cmd", "/C", "ver")
	} else if runtime.GOOS == "linux"{
		cmd = exec.Command("uname", "-a")
	} else if runtime.GOOS == "darwin" {
		cmd = exec.Command("sw_vers")
	} else {
		fmt.Println("SO não suportado!")
		return 
	}
	out, err := cmd.CombinedOutput()
	if err != nil {
		return
	}
	fmt.Println(string(out))
}

// 4 - Configuração da máquina
func getSystemInfo() {
    var cmd *exec.Cmd

    if runtime.GOOS == "windows" {
        cmd = exec.Command("cmd", "/C", "systeminfo")
    } else {
        fmt.Println("SO não suportado!")
        return
    }

    out, err := cmd.CombinedOutput()
    if err != nil {
        fmt.Println("Erro ao executar comando: ", err)
        return
    }

    fmt.Println(string(out))
}


func main() {

	fmt.Println(getCurrentDirectory())
	listFilesAndDicts()
	goOSVersion()
	getSystemInfo()
}