package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string

	fmt.Print("Digite uma string: ")
	fmt.Scanln(&input)

	numPalavras := contarPalavras(input)

	fmt.Printf("A string contém %d palavras.\n", numPalavras)
}

func contarPalavras(str string) int {

  palavras := strings.Fields(str)
	return len(palavras)
}
