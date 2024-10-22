package main //Todo programa em GO começa no package main

import (
	"fmt"
)

func main() { // Precisamos da func main para rodar o programa em GO
	// Palavra var não é necessário
	nome := "Felipe"          // Aqui eu não tipei mas o Go inferiu o tipo string // := declara e atribui o valor a variável
	var version float32 = 1.1 // Aqui está tipado como float32 Sem a tipagem ele pode escolher a tipagem que usa mais memória

	fmt.Println("Olá sr.", nome)
	fmt.Println("Este programa está na versão", version)
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("3 - Sair do Programa")

	var comando int
	// fmt.Scanf("%d", &comando) // o & indica o endereço da variável onde vai ser inserido o valor
	fmt.Scan(&comando) // Não tem necessidade e especificar o tipo recebido com o %

	switch comando {
	case 1:
		fmt.Println("Você selecionou 1")
	case 2:
		fmt.Println("Você selecionou 2")
	case 3:
		fmt.Println("Saindo do programa!")
	default:
		fmt.Println("Não reconheço essa instrução")
	}
}
