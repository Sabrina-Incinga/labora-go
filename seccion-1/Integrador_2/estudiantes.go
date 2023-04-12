package main

import (
	"fmt"
	"reflect"
	"sort"
)

type Estudiante struct {
	Nombre string
	Nota   float64
	Codigo string
}

func crearEstudiante(nombre *string, nota *float64, codigo *string) *Estudiante {
	var estudiante *Estudiante

	if *nombre == "" || *nota < 0 || *codigo == "" {
		estudiante = nil
	} else{
		estudiante = &Estudiante{Nombre: *nombre, Nota: *nota, Codigo: *codigo}
	}

	return estudiante
}

func ordenarEstudiantes(estudiantes []Estudiante, criterio int, orden int) []Estudiante{
	var ordenar func(i, j int) bool

	switch criterio{
		case 1:
			switch orden{
				case 1:
					ordenar = func(i, j int) bool {
						return estudiantes[i].Nombre < estudiantes[j].Nombre
					}
				case 2:
					ordenar = func(i, j int) bool {
						return estudiantes[i].Nombre > estudiantes[j].Nombre
					}
			}
		case 2:
			switch orden{
				case 1:
					ordenar = func(i, j int) bool {
						return estudiantes[i].Nota < estudiantes[j].Nota
					}
				case 2:
					ordenar = func(i, j int) bool {
						return estudiantes[i].Nota > estudiantes[j].Nota
					}
			}
		case 3:
			switch orden{
			case 1:
				ordenar = func(i, j int) bool {
					return estudiantes[i].Codigo < estudiantes[j].Codigo
				}
			case 2:
				ordenar = func(i, j int) bool {
					return estudiantes[i].Codigo > estudiantes[j].Codigo
				}
		}
		default:
			return estudiantes
	}

	sort.Slice(estudiantes, ordenar)

	return estudiantes
}

func buscarEstudiante(estudiantes []Estudiante, criterio int, valor any) *Estudiante{
	var estudianteBuscado Estudiante

	for _, v := range estudiantes {
		tipo := reflect.TypeOf(v)

		for i := 0; i < tipo.NumField(); i++ {
			if i+1 == criterio {
				campo := tipo.Field(i)
				valorCampo := reflect.ValueOf(v).FieldByIndex(campo.Index)

				if valorCampo.Interface() == valor {
					estudianteBuscado = v
				}
			}
		}

		if &estudianteBuscado != nil {
			break
		}
	}

	return &estudianteBuscado
}

func crearEstudiantesPorConsola(nombre string, nota float64, codigo string, estudiante Estudiante, estudiantes *[]Estudiante) {
	fmt.Println("Ingrese el nombre del estudiante:")
	fmt.Scanln(&nombre)
	fmt.Println("Ingrese la nota del estudiante:")
	fmt.Scanln(&nota)
	fmt.Println("Ingrese el código del estudiante:")
	fmt.Scanln(&codigo)
	estudiante = *crearEstudiante(&nombre, &nota, &codigo)

	if &estudiante != nil {
		*estudiantes = append(*estudiantes, estudiante)
		fmt.Printf("*********************\nEstudiante creado con éxito: %v\n*********************", estudiante)
		fmt.Println("*********************\nFin de la solicitud\n*********************")
	} else {
		fmt.Println("Estudiante no creado, alguno de los datos es incorrecto o vacío")
		fmt.Println("*********************\nFin de la solicitud\n*********************")
	}
}

func buscarEstudiantePorConsola(criterio int, nombre string, estudiantes []Estudiante, estudiante Estudiante, nota float64, codigo string) {
	fmt.Println("Ingrese el número del criterio por el que desea buscar:\n\t1.Nombre\n\t2.Nota\n\t3.Codigo")
	fmt.Scanln(&criterio)
	switch criterio {
	case 1:
		fmt.Println("Ingrese el nombre del estudiante:")
		fmt.Scanln(&nombre)
		estudiantePtr := buscarEstudiante(estudiantes, criterio, nombre)
		if estudiantePtr != nil {
			estudiante = *estudiantePtr
			fmt.Printf("Estudiante encontrado: \n%v\n", estudiante)
		} else {
			fmt.Println("Estudiante no encontrado")
		}
	case 2:
		fmt.Println("Ingrese la nota del estudiante:")
		fmt.Scanln(&nota)
		estudiantePtr := buscarEstudiante(estudiantes, criterio, nota)
		if estudiantePtr != nil {
			estudiante = *estudiantePtr
			fmt.Printf("Estudiante encontrado: \n%v\n", estudiante)
		} else {
			fmt.Println("Estudiante no encontrado")
		}
	case 3:
		fmt.Println("Ingrese el código del estudiante:")
		fmt.Scanln(&codigo)
		estudiantePtr := buscarEstudiante(estudiantes, criterio, codigo)
		if estudiantePtr != nil {
			estudiante = *estudiantePtr
			fmt.Printf("Estudiante encontrado: \n%v\n", estudiante)
		} else {
			fmt.Println("Estudiante no encontrado")
		}
	}
}


func main() {
	var accion int 
	var criterio int
	var orden int
	var estudiante Estudiante
	var nombre string
	var nota float64
	var codigo string
	var estudiantes []Estudiante

	for accion != 5{
		fmt.Println("Seleccione el número de la opción elegida:\n\t1.Crear estudiante\n\t2.Mostrar lista de estudiantes\n\t3.Ordenar Lista\n\t4.Buscar estudiante\n\t5.Salir")
		fmt.Scanln(&accion)
		switch accion{
			case 1:
				crearEstudiantesPorConsola(nombre, nota, codigo, estudiante, &estudiantes)
				accion = 0
			case 2:
				fmt.Printf("Lista de estudiantes: \n%v\n", estudiantes)
				fmt.Println("*********************\nFin de la solicitud\n*********************")
				accion = 0
			case 3:
				fmt.Println("Ingrese el número del criterio por el que desea ordenar:\n\t1.Nombre\n\t2.Nota\n\t3.Codigo")
				fmt.Scanln(&criterio)
				fmt.Println("Ingrese el número del orden deseado:\n\t1.Ascendente\n\t2.Descendente")
				fmt.Scanln(&orden)
				ordenarEstudiantes(estudiantes,criterio,orden)
				fmt.Printf("Lista ordenada según los criterios seleccionados:\n\t%v\n", estudiantes)
				fmt.Println("*********************\nFin de la solicitud\n*********************")
				accion = 0
			case 4:
				buscarEstudiantePorConsola(criterio, nombre, estudiantes, estudiante, nota, codigo)
				fmt.Println("*********************\nFin de la solicitud\n*********************")
				accion = 0
			case 5:
				fmt.Println("Gracias, vuelva pronto!")
			default:
				fmt.Println("Opción incorrecta, vuelva a intentar:")
				accion = 0
		}
	}

}



