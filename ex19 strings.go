package main

import (
	"fmt"
)

func main() {
	var input string

	fmt.Print("Digite uma string: ")
	fmt.Scanln(&input)

	fmt.Println("String invertida:", inverterString(input))
}

func inverterString(str string) string {
	runes := []rune(str)
	tamanho := len(runes)

	for i := 0; i < tamanho/2; i++ {
		runes[i], runes[tamanho-1-i] = runes[tamanho-1-i], runes[i]
	}

	return string(runes)
}

