package main

import "fmt"

/*
Una cadena de ADN se representa como una secuencia circular de bases (adenina, timina, citosina y guanina) que es única para cada ser vivo, por ejemplo:

A T F
T	C
A T G

Dicha cadena se puede representar como un arreglo de caracteres recorriendola en sentido horario desde la parte superior izquierda:
A T G C G T A T
Se pide diseñar una clase que represente una secuencia de ADN e incluya un método booleano que nos devuelva true si dos cadenas de ADN coinciden.
MUY IMPORTANTE: La secuencia de ADN es cíclica, por lo que puede comenzar en cualquier posición. Por ejemplo, las dos secuencias siguientes coinciden:
A T G C G T A T
A T A T G C G T
Pregunta existencial: ¿la cantidad de combinaciones de ADN debe ser finita? ¿Cuántas combinaciones diferentes puede haber? ¿Puede ocurrir algún día que se empiezan a repetir??
*/

/*
Ir recorriendo la cadena 1 tomando 1 base a la vez, verifico si existe esa base en la cadena 2 y además, si las bases anterior y posterior en la cadena 1 es igual a las de la cadena 2. Si es así, pusheo la base en análisis a un slice, si el slice resultante es igual al de la cadena 1, entonces puedo asumir que cadena 1 y cadena 2 son iguales.
*/

func compareDNASequence(seq1 []string, seq2 []string) bool{
	var response bool = false

	if len(seq2) != len(seq1){
		return response
	}

	for i := 0; i < len(seq1); i++ {
		match := true
		for j := 0; j < len(seq2); j++ {
			if seq1[(i+j)%len(seq1)] != seq2[j] {
				match = false
                break
			}
		}
		if match {
			response = true
			break
        }
	}

	return response
}

func main(){
	seq1 := []string{"A", "T", "G", "C", "G", "T", "A", "T"}
	seq2 := []string{"A", "T", "G", "C", "G", "T", "A", "T"}
	fmt.Println(compareDNASequence(seq1, seq2))
}