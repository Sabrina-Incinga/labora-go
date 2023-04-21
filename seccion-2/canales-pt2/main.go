package main

import (
	"fmt"
	"labora-go/seccion-2/canales-pt2/multiplicacion"
	"labora-go/seccion-2/canales-pt2/suma"
)


func consumer(acum *int,c <-chan int) {
	for n := range c {
		*acum += n
	}
}


func main() {
	rango := 300000
	valorInicial := 1
	ch := make(chan int)

	go suma.SumarRango(valorInicial, rango/2, ch)
	go suma.SumarRango(valorInicial+rango/2, rango, ch)

	x, y := <-ch, <-ch // receive from c

	fmt.Println(x, y, x+y) 
	
	var matrizA [][]int = [][]int{{1,2,3}, {4,5,6}}
	var matrizB [][]int = [][]int{{1,2}, {1,2}, {2,3}}

	multiplicacion.MultiplicarMatrices(matrizA, matrizB)

}