package main

import (
	"fmt"
	"strings"
)

/*
Realice un algoritmo (o sea en un archivo.go codifiquen un programita) para invertir una cadena de caracteres (la cadena puede ser ingresada por teclado o bien puede ser un valor fijo que ustedes gusten,  BTW a ese valor fijo dentro del código se lo llama valor "hardcodeado"). ¿Que pasaría se compara una cadena con la misma cadena al revés? ¿Que problema estoy resolviendo?
*/

func pushLettersToSlice(s string, wordSlice *[]string){
	for _, v := range s {
		b := func (x rune)  {
			*wordSlice = append(*wordSlice, string(x))
		}
		defer b(v)
	}
	
}

func invertWords(word string) string{
	var slice []string
	pushLettersToSlice(word, &slice)

	return strings.Join(slice, "")

}

func main(){
	fmt.Println(invertWords("cuadrado"))
}