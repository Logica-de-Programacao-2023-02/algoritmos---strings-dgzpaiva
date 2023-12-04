package main

import (
	"fmt"
)

func main() {
	var input string

	fmt.Print("Digite uma string: ")
	fmt.Scanln(&input)

	letrasUnicas := encontrarLetrasUnicas(input)

	fmt.Println("Letras únicas na string:", letrasUnicas)
}

func encontrarLetrasUnicas(str string) string {
	ocorrencias := make(map[rune]int)

	for _, char := range str {
		ocorrencias[char]++
	}

	var letrasUnicas string

	for _, char := range str {
		if ocorrencias[char] == 1 {
			letrasUnicas += string(char)
		}
	}

	return letrasUnicas
}
