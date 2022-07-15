package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

const monitoramentos = 3 //quantidade de monitoramento
const delay = 5          //segundos

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
	fmt.Println("") // pulando linha

	return comando
}

func showMenu() {
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do Programa")
	fmt.Println("") // pulando linha
}

func startMonitoring() {
	fmt.Println("Monitorando...")
	round := 1
	//sites := []string{"https://www.alura.com.br", "https://random-status-code.herokuapp.com", "https://www.uol.com.br"}

	sites := readSites()

	for i := 0; i < monitoramentos; i++ {
		fmt.Println("Monitoramento:", round) //adicionando repetição e tempo
		for i, site := range sites {
			fmt.Println(i, "-", site)
			testSites(site)
		}
		round++
		time.Sleep(delay * time.Second) //pacote time
		fmt.Println("")                 // pulando linha
	}

	fmt.Println("") // pulando linha
}

func testSites(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("O site:", site, "foi carregado com sucesso")
	} else {
		fmt.Println("o site", site, "está com problema, Status Code:", resp.StatusCode)
	}
}

//lendo arquivo de texto:
func readSites() []string {

	var sites []string

	arquivo, err := os.Open("sites.txt")
	//arquivo, err := ioutil.ReadFile("sites.txt") //le o arquivo todo

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	leitor := bufio.NewReader(arquivo)

	for {
		linha, err := leitor.ReadString('\n') //pegando até o \n
		linha = strings.TrimSpace(linha)      //tira o \n e espaço em branco no final das linhas

		sites = append(sites, linha)

		if err == io.EOF {
			break
		}

	}
	arquivo.Close()
	return sites
}
