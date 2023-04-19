package main

import("fmt"
		"labora-go/seccion-2/canales-pt2/suma"
		"time")


func consumer(acum *int,c <-chan int) {
	for n := range c {
		*acum += n
	}
}


func main() {
	rango := 300000
	valorInicial := 1
	ch := make(chan int)
	acum := 0

	go suma.SumarRango(valorInicial, rango/2, ch)
	go suma.SumarRango(valorInicial+rango/2, rango, ch)

	go consumer(&acum ,ch)

	
	time.Sleep(time.Second)
	fmt.Println(acum)
}