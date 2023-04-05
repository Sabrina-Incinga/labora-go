package calculadora

import "fmt"

func Calculadora(){
	var a, b int

	fmt.Println("Calculadora")
	fmt.Println("Ingrese el primer número entero:")
	fmt.Scanln(&a)
	fmt.Println("Ingrese el segundo número entero:")
	fmt.Scanln(&b)
	
	fmt.Printf("Suma: %d\n", sumar(a, b))
	fmt.Printf("Resta: %d\n", restar(a, b))
	fmt.Printf("Multiplicación: %d\n", multiplicar(a, b))
	fmt.Printf("División: %d\n", dividir(a, b))

}

func sumar(a int, b int) int{
	return a + b
}

func restar(a int, b int) int{
	return a - b
}

func multiplicar(a int, b int) int{
	return a * b
}

func dividir(a int, b int) int{
	if(b == 0){
		return 0
	} else{
		return a / b
	}
}