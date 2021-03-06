package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitoring = 3
const timeDelay = 2

func main() {

	for {
		introduction()
		menu()
		numberComand := getComand()

		switch numberComand {
		case 1:
			initMonitoring()
		case 2:
			fmt.Println("Exibindo Logs...")
			getLogs()
		case 0:
			fmt.Println("Saindo do programa...")
			os.Exit(0)
		default:
			fmt.Println("Não conheço este comando")
			os.Exit(-1)
		}

	}
}

func introduction() {
	name := "Tony"
	version := 1.1

	fmt.Println("Olá", name)
	fmt.Println("Bem-vindo ao monitoramento, esse programa está na versão", version)
}

func menu() {
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do programa")
}

func getComand() int {
	var comand int
	fmt.Scan(&comand)
	fmt.Println("")
	return comand
}

func initMonitoring() {
	fmt.Println("Monitorando...")
	// sites := []string{"https://slack.com/intl/pt-br/", "https://github.com/", "https://hub.docker.com/", "https://www.linkedin.com/"}
	sites := getArchiveSites()
	for i := 0; i < monitoring; i++ {
		for i, site := range sites {
			fmt.Println("Testando site", i, ":", site)
			siteTest(site)
		}
		time.Sleep(timeDelay * time.Second)
		fmt.Println("")
	}
}

func siteTest(site string) {
	response, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro", nil)
	}

	if response.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
		logRegister(site, true)
	} else {
		fmt.Println("Site:", site, "esta com problemas. Status Code:", response.StatusCode)
		logRegister(site, false)
	}
}

func getArchiveSites() []string {
	// archive, err := ioutil.ReadFile("sites.txt")
	var sites []string

	archive, err := os.Open("sites.txt")
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	reader := bufio.NewReader(archive)
	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		sites = append(sites, line)
		if err == io.EOF {
			break
		}
	}

	archive.Close()

	return sites
}

func logRegister(site string, status bool) {

	archive, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
	}

	archive.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online: " + strconv.FormatBool(status) + "\n")

	archive.Close()
}

func getLogs() {
	archive, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(archive))
}
