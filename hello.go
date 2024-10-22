package main //Todo programa em GO começa no package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() { // Precisamos da func main para rodar o programa em GO
	for { // Go não tem while. O for dessa forma seria a mesma coisa

		// Palavra var não é necessário
		nome := "Felipe"          // Aqui eu não tipei mas o Go inferiu o tipo string // := declara e atribui o valor a variável
		var version float32 = 1.1 // Aqui está tipado como float32 Sem a tipagem ele pode escolher a tipagem que usa mais memória

		saudacao(nome, version)
		menu()

		comando := leComando()
		// fmt.Scanf("%d", &comando) // o & indica o endereço da variável onde vai ser inserido o valor

		switch comando {
		case 1:
			monitoramento()
		case 2:
			fmt.Println("Exibir logs")
		case 3:
			fmt.Println("Saíndo do programa!")
			os.Exit(0) // Terminando o programa com sucesso! caso erro podemos contar com um retorno -1
		default:
			fmt.Println("Não reconheço essa instrução!")
			os.Exit(-1)
		}
	}
}

func saudacao(nome string, version float32) {
	fmt.Println("Olá sr.", nome)
	fmt.Println("Este programa está na versão", version)
}

func menu() {
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("3 - Sair do Programa")
}

func leComando() int {
	var comando int
	fmt.Scan(&comando) // Não tem necessidade e especificar o tipo recebido com o %
	return comando
}

func monitoramento() {
	fmt.Println("Monitorando...")
	sites := []string{
		"https://go.dev",
		"https://www.infomoney.com.br"}

	for i := 0; i < len(sites); i++ {
		fmt.Println(sites[i])
		testaSite(sites[i])
	}
}

func testaSite(site string) {
	resp, _ := http.Get(site)
	// if err != nil { // tratando erro caso exista
	// 	fmt.Println("Hmm... Site não parece estar ok! Status code:")
	// }

	if resp.StatusCode == 200 {
		fmt.Println("Site está rodando", resp.StatusCode, "Em: ")
	} else {
		fmt.Println("Site", site, " não está rodando...")
	}
}

func Logs() {}
