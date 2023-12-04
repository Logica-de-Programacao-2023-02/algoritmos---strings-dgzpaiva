package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	var str1, str2 string

	fmt.Print("Digite a primeira string: ")
	fmt.Scanln(&str1)

	fmt.Print("Digite a segunda string: ")
	fmt.Scanln(&str2)

	saoAnagramas := verificarAnagramas(str1, str2)

	if saoAnagramas {
		fmt.Println("As strings são anagramas.")
	} else {
		fmt.Println("As strings não são anagramas.")
	}
}

func verificarAnagramas(str1, str2 string) bool {

	str1 = strings.ToLower(str1)
	str2 = strings.ToLower(str2)

	runes1 := []rune(str1)
	runes2 := []rune(str2)

	sort.Slice(runes1, func(i, j int) bool { return runes1[i] < runes1[j] })
	sort.Slice(runes2, func(i, j int) bool { return runes2[i] < runes2[j] })

	return string(runes1) == string(runes2)
}
