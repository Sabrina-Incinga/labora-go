package multiplicacion

import "fmt"

/*
**Ejercicio 2**: Multiplicación de matrices utilizando gorutinas y canales
Dadas dos matrices A y B, crea un programa en Go que realice la multiplicación de las matrices utilizando gorutinas y canales para dividir el trabajo entre varias gorutinas. Las matrices se pueden representar como matrices bidimensionales (slices) en Go.
*/

//len == nro filas 
//len de elementos == nro de columnas 
//len x lenElementos
type FilaResultado struct {
    Indice int
    Fila   []int
}

// valida si una matriz es rectangular, es decir, que todas sus filas tienen el mismo número de columnas, devuelve un booleano y el número de columnas en caso de ser matriz rectangular
func esMatrizRectangular(matriz [][]int) (int, bool) {
	var nroColumnas int
	var valid bool
	if len(matriz) != 0 {
		nroColumnas = len(matriz[0])
		for _, v := range matriz {
			if nroColumnas == len(v) {
				valid = true
			} else {
				valid = false
			}
		}
	}

	if !valid {
		return 0, valid
	}
	return nroColumnas, valid
}

// multiplica una fila de la matriz A por la matriz B y envía el resultado por el canal result
func multiplicarFila(A [][]int, fila int, B [][]int, resultado chan<- FilaResultado) {
	filaA := A[fila]
	filaResultado := make([]int, len(B[0]))
	for i := 0; i < len(B[0]); i++ {
		colB := make([]int, len(B))
		for j := 0; j < len(B); j++ {
			colB[j] = B[j][i]
		}
		filaResultado[i] = productoPunto(filaA, colB)
	}
	resultado <- FilaResultado{Indice: fila, Fila: filaResultado}
}

// realiza el producto punto entre dos vectores
func productoPunto(v1 []int, v2 []int) int {
	resultado := 0
	for i := 0; i < len(v1); i++ {
		resultado += v1[i] * v2[i]
	}
	return resultado
}

// imprime una matriz en la consola
func imprimirMatriz(matriz [][]int) {
	for i := 0; i < len(matriz); i++ {
		for j := 0; j < len(matriz[0]); j++ {
			fmt.Printf("[%d][%d]: %d ",i, j, matriz[i][j])
		}
		fmt.Println()
	}
	fmt.Println()
}

func MultiplicarMatrices(matrizA [][]int, matrizB [][]int){
	//validamos que sean matrices rectangulares, es decir, matrices válidas para multiplicar
	nroColumnasA, esRectangularA := esMatrizRectangular(matrizA)
	_, esRectangularB := esMatrizRectangular(matrizB)
	nroFilasB := len(matrizB)
	matrizResultado := make([][]int, len(matrizA))
	
	//validamos que las matrices válidas puedan multiplicarse, verificando 	que el nro de columnas de matrizA sea igual al nro de filas de la matriB
	if esRectangularA && esRectangularB {
		if nroColumnasA == nroFilasB {
			ch := make(chan FilaResultado, len(matrizA))
			for i := 0; i < len(matrizA); i++ {
				go multiplicarFila(matrizA, i, matrizB, ch)
			}

			for i := 0; i < len(matrizA); i++ {
				rowResult := <-ch
				matrizResultado[rowResult.Indice] = rowResult.Fila
			}
		}
	}

	imprimirMatriz(matrizResultado)
	//PROBLEMA: no es posible predecir qué go rutina se ejecutará primero, por lo que la matriz resultante no necesariamente se forma en el orden correcto
	//SOLUCIÓN: dado que el número de las filas de la matriz solución coincide con el de la matriz a, se envía por el canal de resultado tanto el número de fila, que enviamos por parámetro en la función multiplyRow, como la fila resultante 
}