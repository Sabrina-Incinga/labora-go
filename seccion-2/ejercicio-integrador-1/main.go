package main

import(
	"fmt"
	persona "labora-go/seccion-2/ejercicio-integrador-1/pp2"
	"labora-go/seccion-2/ejercicio-integrador-1/pp4"
)

func main() {
	var nombre string
	var edad int
	var altura float64
	var peso float64
	var accion int
	var personas []persona.Persona
	var criterio int

	defer func() {
		pp4.CalcularPesoProm(&personas)
		pp4.MinyMax(personas)
	}()

	for accion != 5{
		fmt.Println("Ingrese el número de la opción deseada: \n1. Crear persona \n2. Ordenar personas \n3. Buscar persona \n4. Calcular IMC de una persona \n5. Salir")
		fmt.Scanln(&accion)
		switch accion{
		case 1:
			fmt.Println("Ingrese el nombre: ")
			fmt.Scanln(&nombre)
			fmt.Println("Ingrese la edad: ")
			fmt.Scanln(&edad)
			fmt.Println("Ingrese la altura: ")
			fmt.Scanln(&altura)
			fmt.Println("Ingrese el peso: ")
			fmt.Scanln(&peso)

			persona := persona.CrearPersona(&nombre, &edad, &altura, &peso)
			personas = append(personas, *persona)
			fmt.Printf("********************** \nPersona creada exitosamente: \n%v\n**********************\n", *persona)
			fmt.Println("*********************************************Fin Solicitud**********************************************")
			accion = 0
		case 2:
			fmt.Println("Ingrese el número de propiedad por la que desea ordenar: \n1. Nombre \n2. Edad \n3. Altura \n4. Peso ")
			fmt.Scanln(&criterio)
			fmt.Println("Personas sin ordenar")
			fmt.Println(personas)
			fmt.Println("Personas ordenadas según el criterio elegido:")
			persona.OrdenarPersonas(personas, criterio)
			fmt.Println(personas)
			fmt.Println("*********************************************Fin Solicitud**********************************************")
			accion = 0
		case 3:
			persona.BuscarPersonaPorConsola(&criterio, &nombre, &personas, &edad, &altura, &peso)
			accion = 0
		case 4:
			persona.CalcularIMCPorConsola(&criterio, &nombre, &personas, &edad, &altura, &peso)
			accion = 0
		case 5:
			fmt.Println("Gracias, vuelva pronto!")
		default:
			accion = 0
		}
	}

}