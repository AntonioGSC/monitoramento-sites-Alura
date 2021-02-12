package main

import "fmt"

func main() {
	name := "Tony"
	version := 1.1

	fmt.Println("Olá", name)
	fmt.Println("Bem-vindo ao monitoramento, esse programa está na versão", version)

	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do programa")

	var number int
	//fmt.Scanf("%d", &number)
	fmt.Scan(&number)
	fmt.Println("Variavel a seguir", number)
}
