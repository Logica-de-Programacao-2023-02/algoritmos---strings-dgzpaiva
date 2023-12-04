package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string

	fmt.Print("Digite uma string: ")
	fmt.Scanln(&input)

	ePalindromo := verificarPalindromo(input)

	if ePalindromo {
		fmt.Println("A string é um palíndromo.")
	} else {
		fmt.Println("A string não é um palíndromo.")
	}
}

func verificarPalindromo(str string) bool {

	str = strings.ReplaceAll(strings.ToLower(str), " ", "")

	return str == inverterString(str)
}

func inverterString(str string) string {
	runes := []rune(str)
	tamanho := len(runes)

	for i := 0; i < tamanho/2; i++ {
		runes[i], runes[tamanho-1-i] = runes[tamanho-1-i], runes[i]
	}

	return string(runes)
}
