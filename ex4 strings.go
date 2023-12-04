package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string
	var antigo, novo rune

	fmt.Print("Digite uma string: ")
	fmt.Scanln(&input)

	fmt.Print("Digite o caractere a ser substituído: ")
	fmt.Scanf("%c", &antigo)

	fmt.Print("Digite o novo caractere: ")
	fmt.Scanf(" %c", &novo)

	resultado := substituirCaractere(input, antigo, novo)

	fmt.Println("String com substituição:", resultado)
}

func substituirCaractere(str string, antigo, novo rune) string {
	runes := []rune(str)
	for i, char := range runes {
		if char == antigo {
			runes[i] = novo
		}
	}

	return string(runes)
}
