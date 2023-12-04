package main

import (
	"fmt"
	"strings"
)

func main() {
	var str1, str2 string

	fmt.Print("Digite a primeira string: ")
	fmt.Scanln(&str1)

	fmt.Print("Digite a segunda string: ")
	fmt.Scanln(&str2)

	eSubstring := verificarSubstring(str1, str2)

	if eSubstring {
		fmt.Println("A segunda string é uma substring da primeira.")
	} else {
		fmt.Println("A segunda string NÃO é uma substring da primeira.")
	}
}

func verificarSubstring(str1, str2 string) bool {

	return strings.Contains(str1, str2)
}
