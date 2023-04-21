package multiplicacion

import "fmt"

/*
**Ejercicio 2**: Multiplicación de matrices utilizando gorutinas y canales
Dadas dos matrices A y B, crea un programa en Go que realice la multiplicación de las matrices utilizando gorutinas y canales para dividir el trabajo entre varias gorutinas. Las matrices se pueden representar como matrices bidimensionales (slices) en Go.
*/

//len == nro filas 
//len de elementos == nro de columnas 
//len x lenElementos

func MultiplicarMatrices(matrizA [][]int, matrizB [][]int){
	//validamos que sean matrices rectangulares, es decir, matrices válidas para multiplicar
	var validA, validB bool
	var nroColumnasA int = 0
	matrizResultado := make([][]int, len(matrizA))
	if len(matrizA) != 0 {
		nroColumnasA = len(matrizA[0])
		for _, v := range matrizA {
			if nroColumnasA == len(v) {
				validA = true
			}else {
				validA = false
			}
		}
	}
	
	nroFilasB := len(matrizB)
	if nroFilasB != 0 {
		nroColumnasB := len(matrizB[0])
		for _, v := range matrizB {
			if nroColumnasB == len(v) {
				validB = true
			}else {
				validB = false
			}
		}
	}
	//validamos que las matrices válidas puedan multiplicarse, verificando 	que el nro de columnas de matrizA sea igual al nro de filas de la matriB
	if validA && validB {
		if nroColumnasA == nroFilasB {
			var ch chan []int = make(chan []int)
			for i := 0; i < len(matrizA); i++ {
				go multiplyRow(matrizA, i, matrizB, ch)
			}

			for i := 0; i < len(matrizA); i++ {
				matrizResultado[i] = <-ch
			}
		}
	}

	printMatrix(matrizResultado)
	//PROBLEMA: no es posible predecir qué go rutina se ejecutará primero, por lo que la matriz resultante no necesariamente se forma en el orden correcto
}

// multiplica una fila de la matriz A por la matriz B y envía el resultado por el canal result
func multiplyRow(A [][]int, row int, B [][]int, result chan<- []int) {
	rowA := A[row]
	rowResult := make([]int, len(B[0]))
	for i := 0; i < len(B[0]); i++ {
		colB := make([]int, len(B))
		for j := 0; j < len(B); j++ {
			colB[j] = B[j][i]
		}
		rowResult[i] = dotProduct(rowA, colB)
	}
	result <- rowResult
}

// realiza el producto punto entre dos vectores
func dotProduct(v1 []int, v2 []int) int {
	result := 0
	for i := 0; i < len(v1); i++ {
		result += v1[i] * v2[i]
	}
	return result
}

// imprime una matriz en la consola
func printMatrix(matrix [][]int) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			fmt.Printf("matrix[%d][%d]: %d ",i, j, matrix[i][j])
		}
		fmt.Println()
	}
	fmt.Println()
}