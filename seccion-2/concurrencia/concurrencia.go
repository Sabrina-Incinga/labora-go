 package main 

 import (
	"fmt"
	"sync"
	"time"
)

func sumarSecuencial(numeros []int, acum *int){
	for _, v := range numeros {
		*acum += v
	}

}

func sumar(numeros []int, acum *int, wg *sync.WaitGroup){
	defer wg.Done()
	for _, v := range numeros {
		*acum += v
	}
}

func main(){
	var wg sync.WaitGroup
	var numeros []int
	limit := 1000000

	for i := 0; i <= limit; i++ {
		numeros = append(numeros, i)
	}
	
	wg.Add(2)
	inicio := time.Now()
	var suma1 int
	var suma2 int
	var subLimite int = limit/2

	go sumar(numeros[:subLimite], &suma1, &wg)

	go sumar(numeros[subLimite:], &suma2, &wg)

	wg.Wait()
	sumaTotal := suma1 + suma2
	duracion := time.Since(inicio)
	fmt.Println("duración con go routines: ", duracion)
	fmt.Println("Suma total: ", sumaTotal)

	inicioSecuencial := time.Now()
	var sumaSecuencial int
	sumarSecuencial(numeros, &sumaSecuencial)

	duracionSecuencial := time.Since(inicioSecuencial)

	fmt.Println("duración suma secuencial: ", duracionSecuencial)
	fmt.Println("Suma total secuencial: ", sumaSecuencial)
}
