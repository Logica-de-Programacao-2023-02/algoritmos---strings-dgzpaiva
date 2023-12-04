package main

import (
	"fmt"
	"unicode"
)

func main() {
	var input string

	fmt.Print("Digite uma string: ")
	fmt.Scanln(&input)

	// Verificar se a string está em camelCase e contar as palavras
	eCamelCase, numPalavras := verificarCamelCase(input)

	// Exibir o resultado para o usuário
	if eCamelCase {
		fmt.Println("A string está em camelCase.")
	} else {
		fmt.Println("A string NÃO está em camelCase.")
	}

	fmt.Printf("A string possui %d palavras.\n", numPalavras)
}

func verificarCamelCase(str string) (bool, int) {

	if !unicode.IsLower(rune(str[0])) {
		return false, 0
	}

	var numPalavras int
	tamanho := len(str)

	for i := 0; i < tamanho; i++ {

		if unicode.IsUpper(rune(str[i])) && i > 0 {
			numPalavras++
		}
	}

	numPalavras++

	for i := tamanho - 1; i >= 0; i-- {
		if unicode.IsLower(rune(str[i])) {
			return false, numPalavras
		}
	}

	return true, numPalavras
}
