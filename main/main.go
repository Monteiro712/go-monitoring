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

const monitoramento = 2
const delay = 5

func main() {

    for {
        exibeMenu()
        comando := leComando()

        switch comando {
        case 1:
            iniciarMonitoramento()
        case 2:
            fmt.Println("Exibindo Logs...")
			imprimeLogs()
        case 0:
            fmt.Println("Saindo do programa")
            os.Exit(0)
        default:
            fmt.Println("Não conheço este comando")
            os.Exit(-1)
        }
    }
}

func exibeMenu() {
	fmt.Println("1-iniciar monitoramento")
	fmt.Println("2-exibindo logs")
	fmt.Println("0-sair do programa")
}

func leComando() int{
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("o comando lido foi:", comandoLido)

	return comandoLido
}

func iniciarMonitoramento() {
    fmt.Println("Monitorando...")
	sites := leSitesDoArquivo() 

	for i := 0; i < monitoramento; i++ {
		for i, site := range sites {
			fmt.Println("testando o site", i, ":",  site)
			testaSite(site)
		}
		time.Sleep(delay * time.Second)
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
		fmt.Println("o site:", site, "funciona.")
		registraLog(site, true)
	} else {
		fmt.Println("o site:", site, "não carregou. Status code:", resp.StatusCode)
		registraLog(site, false)
	}
}

func leSitesDoArquivo() []string {
	var sites []string
	arquivo, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("ocorreu erro:", err)
	}

	leitor := bufio.NewReader(arquivo)

	for{
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)

		sites = append(sites, linha)
		//end of file
		if err == io.EOF {
			break
		}
	}
	arquivo.Close()
	return sites
}

func registraLog(site string, status bool) {
	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online " + strconv.FormatBool(status) + "\n")

	arquivo.Close()
}

func imprimeLogs() {

	arquivo, err := os.ReadFile("log.txt")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(arquivo))

}

