package main

import (
	"fmt"
	"os"
)

func main() {
	introduction()
	menu()
	numberComand := getComand()

	switch numberComand {
	case 1:
		fmt.Println("Monitorando...")
	case 2:
		fmt.Println("Exibindo Logs...")
	case 0:
		fmt.Println("Saindo do programa...")
		os.Exit(0)
	default:
		fmt.Println("Não conheço este comando")
		os.Exit(-1)
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
	return comand
}
