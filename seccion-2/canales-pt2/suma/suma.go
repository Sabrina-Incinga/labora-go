package suma

/*
**Ejercicio 1**: Suma de números en un rango utilizando gorutinas y canales
Escribe un programa en Go que sume todos los números en un rango dado (por ejemplo, de 1 a 100) utilizando gorutinas y canales para dividir el trabajo entre varias gorutinas.
*/

func SumarRango(valorInicial int, rango int, ch chan<- int){
	var acum int = 0
	for i := valorInicial; i <= rango; i++ {
		acum += i
	}
	ch <- acum
}