package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string
	var letraAntiga, letraNova string

	fmt.Print("Digite uma string: ")
	fmt.Scanln(&input)

	fmt.Print("Digite a letra a ser substituída: ")
	fmt.Scanln(&letraAntiga)

	fmt.Print("Digite a nova letra: ")
	fmt.Scanln(&letraNova)

	resultado := substituirLetra(input, letraAntiga, letraNova)

	fmt.Println("String com substituição:", resultado)
}

func substituirLetra(str, antiga, nova string) string {

	novaString := strings.ReplaceAll(str, antiga, nova)

	return novaString
}
