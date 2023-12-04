package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string

	fmt.Print("Digite uma string: ")
	fmt.Scanln(&input)
	resultado := removerEspacos(input)

	fmt.Println("String sem espaços em branco:", resultado)
}

func removerEspacos(str string) string {
	semEspacos := strings.ReplaceAll(str, " ", "")

	return semEspacos
}
