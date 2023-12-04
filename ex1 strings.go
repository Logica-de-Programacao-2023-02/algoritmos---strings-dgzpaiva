package main

import (
	"fmt"
)

func main() {
	var input string

	fmt.Print("Digite uma string: ")
	fmt.Scanln(&input)

	resultado := converterParaMaiusculas(input)

	fmt.Println("String em maiúsculas:", resultado)
}

func converterParaMaiusculas(str string) string {
	// Variável para armazenar a string convertida
	var converted string

	for _, char := range str {
		if char >= 'a' && char <= 'z' {
			converted += string(char - 32)
		} else {
			converted += string(char)
		}
	}

	return converted
}
