package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string

	fmt.Print("Digite uma string: ")
	fmt.Scanln(&input)

	resultado := substituirVogais(input)

	fmt.Println("String com vogais substituídas por asteriscos:", resultado)
}

func substituirVogais(str string) string {

	str = strings.ReplaceAll(str, "a", "*")
	str = strings.ReplaceAll(str, "e", "*")
	str = strings.ReplaceAll(str, "i", "*")
	str = strings.ReplaceAll(str, "o", "*")
	str = strings.ReplaceAll(str, "u", "*")
	str = strings.ReplaceAll(str, "A", "*")
	str = strings.ReplaceAll(str, "E", "*")
	str = strings.ReplaceAll(str, "I", "*")
	str = strings.ReplaceAll(str, "O", "*")
	str = strings.ReplaceAll(str, "U", "*")

	return str
}
