package main

import (
	"fmt"
	"unicode"
)

func main() {
	var input string

	fmt.Print("Digite uma string: ")
	fmt.Scanln(&input)

	temNumero := verificarNumero(input)

	if temNumero {
		fmt.Println("A string contém pelo menos um número.")
	} else {
		fmt.Println("A string não contém nenhum número.")
	}
}

func verificarNumero(str string) bool {

	for _, char := range str {
		if unicode.IsDigit(char) {
			return true
		}
	}
	return false
}
