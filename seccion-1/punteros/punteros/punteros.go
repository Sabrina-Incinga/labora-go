package punteros
import "fmt"

func Punteros(){
	var a, b = 10, 20
	var ptrA *int = &a

	fmt.Println("Punteros")
	fmt.Println("Valores originales")
	fmt.Printf("a = %d, b = %d\n", a, b)
	
	*ptrA, b = b, *ptrA
	fmt.Println("Valores intercambiados")
	fmt.Printf("a = %d, b = %d\n", a, b)
}