package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitoramentos = 3
const delay = 5

func main() {
	exibeIntroducao()

	for {
		exibeMenu()

		comando := leComando()

		// if comando == 1 {
		// 	fmt.Println("Monitorando...")
		// } else if comando == 2 {
		// 	fmt.Println("Exibindo logs...")
		// } else if comando == 0 {
		// 	fmt.Println("Saindo do programa...")
		// } else {
		// 	fmt.Println("Comando não reconhecido")
		// }

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo logs...")
			imprimeLogs()
		case 0:
			fmt.Println("Saindo do programa...")
			os.Exit(0)
		default:
			fmt.Println("Comando não reconhecido")
			os.Exit(-1)
		}
	}
}

func exibeIntroducao() {
	// var nome string = "Marcio Coelho"
	// var idade int = 49
	// var versao float32 = 1.1

	// var nome = "Marcio Coelho"
	// var idade = 49
	// var versao float32 = 1.1

	nome := "Marcio Coelho"
	//idade := 49
	versao := 1.1

	fmt.Println("Olá sr.", nome)
	fmt.Println("Este programa está na versão", versao)
}

func exibeMenu() {
	fmt.Println("1. Iniciar monitoramento")
	fmt.Println("2. Exibir logs")
	fmt.Println("0. Sair")
}

func leComando() int {
	var comandoLido int
	//fmt.Scanf("%d", &comandoLido)
	fmt.Scan(&comandoLido)
	fmt.Println("")

	//fmt.Println("O comando escolhido foi", comandoLido)
	return comandoLido
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")

	// site com URL inexistente
	// site := "https://httpbin.org/status/200" // 200 o 404
	// resp, _ := http.Get(site)

	// if resp.StatusCode == 200 {
	// 	fmt.Println("Site:", site, "foi carregado com sucesso!")
	// } else {
	// 	fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
	// }

	// // ARRAY
	// var sites [4]string
	// sites[0] = "https://httpbin.org/status/200"
	// sites[1] = "https://httpbin.org/status/404"
	// sites[2] = "https://httpbin.org/status/200"
	// sites[3] = "https://httpbin.org/status/404"
	// fmt.Println(sites)
	// fmt.Println("O array Sites tem tamanho de", len(sites))
	// fmt.Println("O array Sites tem capacidade de", cap(sites))

	// // SLICE - abstração de array
	// sliceSites := []string{"https://httpbin.org/status/200", "https://httpbin.org/status/404", "https://httpbin.org/status/200", "https://httpbin.org/status/404"}
	// sliceSites = append(sliceSites, "https://www.marciocoelho.com.br")
	// fmt.Println(sliceSites)
	// fmt.Println("O sliceSites tem tamanho de", len(sliceSites))
	// fmt.Println("O sliceSites tem capacidade de", cap(sliceSites))

	// sites := []string{"https://httpbin.org/status/200", "https://httpbin.org/status/404",
	// 	"https://httpbin.org/status/200", "https://httpbin.org/status/404", "https://www.marciocoelho.com.br"}

	sites := leSitesDoArquivo()

	for h := 0; h < monitoramentos; h++ {

		//for i := 0; i < len(sites); i++ {
		for i, site := range sites {
			fmt.Println("Testando site", i, ":", site)
			testaSite(site)
		}

		time.Sleep((delay * time.Second))
		fmt.Println("")

	}
	fmt.Println("")

}

func testaSite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
		registraLog(site, true)
	} else {
		fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
		registraLog(site, false)
	}
}

func leSitesDoArquivo() []string {
	var sites []string

	arquivo, err := os.Open("sites.txt")
	//arquivo, err := os.ReadFile("sites.txt")
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	//fmt.Println(string(arquivo))
	leitor := bufio.NewReader(arquivo)

	// linha, err := leitor.ReadString('\n')
	// if err != nil {
	// 	fmt.Println("Ocorreu um erro:", err)
	// }
	// fmt.Println(linha)

	for {

		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)

		sites = append(sites, linha)

		if err == io.EOF {
			break
		}
	}

	//fmt.Println(sites)

	arquivo.Close()

	return sites
}

func registraLog(site string, status bool) {

	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online: " + strconv.FormatBool(status) + "\n")

	arquivo.Close()
}

func imprimeLogs() {

	arquivo, err := os.ReadFile("log.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	fmt.Println(string(arquivo))
}
