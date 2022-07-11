package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	showIntro()

	for {
		showMenu()
		comando := readComando()

		//if e else
		if comando == 1 {
			startMonitoring()
		} else if comando == 2 {
			fmt.Println("Exibindo logs")
		} else if comando == 0 {
			fmt.Println("Saindo do programa")
			os.Exit(0)
		} else {
			fmt.Println("Comando não reconhecido")
			os.Exit(-1)
		}
	}

	//switch
	//switch comando {
	//case 1:
	//	fmt.Println("Monitorando...")
	//case 2:
	//	fmt.Println("Monitorando 2")
	//case 0:
	//	fmt.Println("Monitorando 0")
	//default:
	//	fmt.Println("Comando não reconhecido")
	//}

	//função

}

func showIntro() {
	nome := "Bruno"
	version := 1.0
	fmt.Println("Bem vindo", nome)
	fmt.Println("Esse programa está na versão:", version)
}

func readComando() int {
	var comando int
	//fmt.Scanf("%d", &comando)
	fmt.Scan(&comando)
	fmt.Println("O comando escolhido foi:", comando)

	return comando
}

func showMenu() {
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do Programa")
}

func startMonitoring() {
	fmt.Println("Monitorando...")
	site1 := "https://www.alura.com.br"
	resp, _ := http.Get(site1)

	if resp.StatusCode == 200 {
		fmt.Println("O site:", site1, "foi carregado com sucesso")
	} else {
		fmt.Println("o site", site1, "está com problema, Status Code:", resp.StatusCode)
	}

}
