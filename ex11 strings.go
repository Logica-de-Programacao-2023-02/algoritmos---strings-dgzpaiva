package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string

	fmt.Print("Digite uma string: ")
	fmt.Scanln(&input)

	resultado := removerVogais(input)

	fmt.Println("String sem vogais:", resultado)
}

func removerVogais(str string) string {

	semVogais := strings.Map(func(r rune) rune {
		switch r {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
			return -1 
		default:
			return r
		}
	}, str)

	return semVogais
}
