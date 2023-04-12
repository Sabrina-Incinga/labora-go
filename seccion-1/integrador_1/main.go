package main

import ( "fmt"
		"sort"
		"reflect")

type Persona struct {
	Nombre string
	Edad   int
	Altura float64
	Peso   float64
}

func calcularIMCPorConsola(criterio *int, nombre *string, personas *[]Persona, edad *int, altura *float64, peso *float64) {
	fmt.Println("Debe seleccionar una persona para calcular su IMC \nIngrese el número de propiedad por la que desea buscar: \n1. Nombre \n2. Edad \n3. Altura \n4. Peso ")
	fmt.Scanln(criterio)
	switch *criterio {
	case 1:
		fmt.Println("Ingrese el nombre de la persona buscada:")
		fmt.Scanln(nombre)
		personaBucada := buscarPersona(*personas, *criterio, *nombre)
		if personaBucada == nil {
			fmt.Println("Persona no encontrada")
			fmt.Println("*********************************************Fin Solicitud**********************************************")
		} else {
			fmt.Println("Persona encontrada:")
			fmt.Println(calcularIMC(*personaBucada))
			fmt.Println("*********************************************Fin Solicitud**********************************************")
		}
	case 2:
		fmt.Println("Ingrese la edad de la persona buscada:")
		fmt.Scanln(edad)
		personaBucada := buscarPersona(*personas, *criterio, *edad)
		if personaBucada == nil {
			fmt.Println("Persona no encontrada")
			fmt.Println("*********************************************Fin Solicitud**********************************************")
		} else {
			fmt.Println("Persona encontrada:")
			fmt.Println(calcularIMC(*personaBucada))
			fmt.Println("*********************************************Fin Solicitud**********************************************")
		}
	case 3:
		fmt.Println("Ingrese la altura de la persona buscada:")
		fmt.Scanln(altura)
		personaBucada := buscarPersona(*personas, *criterio, *altura)
		if personaBucada == nil {
			fmt.Println("Persona no encontrada")
			fmt.Println("*********************************************Fin Solicitud**********************************************")
		} else {
			fmt.Println("Persona encontrada:")
			fmt.Println(calcularIMC(*personaBucada))
			fmt.Println("*********************************************Fin Solicitud**********************************************")
		}
	case 4:
		fmt.Println("Ingrese el peso de la persona buscada:")
		fmt.Scanln(peso)
		personaBucada := buscarPersona(*personas, *criterio, *peso)
		if personaBucada == nil {
			fmt.Println("Persona no encontrada")
			fmt.Println("*********************************************Fin Solicitud**********************************************")
		} else {
			fmt.Println("Persona encontrada:")
			fmt.Println(calcularIMC(*personaBucada))
			fmt.Println("*********************************************Fin Solicitud**********************************************")
		}
	}
}

func buscarPersonaPorConsola(criterio *int, nombre *string, personas *[]Persona, edad *int, altura *float64, peso *float64) {
	fmt.Println("Ingrese el número de propiedad por la que desea buscar: \n1. Nombre \n2. Edad \n3. Altura \n4. Peso ")
	fmt.Scanln(criterio)
	switch *criterio {
	case 1:
		fmt.Println("Ingrese el nombre de la persona buscada:")
		fmt.Scanln(nombre)
		personaBucada := buscarPersona(*personas, *criterio, *nombre)
		if personaBucada == nil {
			fmt.Println("Persona no encontrada")
			fmt.Println("*********************************************Fin Solicitud**********************************************")
		} else {
			fmt.Println("Persona encontrada:")
			fmt.Println(*personaBucada)
			fmt.Println("*********************************************Fin Solicitud**********************************************")
		}
	case 2:
		fmt.Println("Ingrese la edad de la persona buscada:")
		fmt.Scanln(edad)
		personaBucada := buscarPersona(*personas, *criterio, *edad)
		if personaBucada == nil {
			fmt.Println("Persona no encontrada")
			fmt.Println("*********************************************Fin Solicitud**********************************************")
		} else {
			fmt.Println("Persona encontrada:")
			fmt.Println(*personaBucada)
			fmt.Println("*********************************************Fin Solicitud**********************************************")
		}
	case 3:
		fmt.Println("Ingrese la altura de la persona buscada:")
		fmt.Scanln(altura)
		personaBucada := buscarPersona(*personas, *criterio, *altura)
		if personaBucada == nil {
			fmt.Println("Persona no encontrada")
			fmt.Println("*********************************************Fin Solicitud**********************************************")
		} else {
			fmt.Println("Persona encontrada:")
			fmt.Println(*personaBucada)
			fmt.Println("*********************************************Fin Solicitud**********************************************")
		}
	case 4:
		fmt.Println("Ingrese el peso de la persona buscada:")
		fmt.Scanln(peso)
		personaBucada := buscarPersona(*personas, *criterio, *peso)
		if personaBucada == nil {
			fmt.Println("Persona no encontrada")
			fmt.Println("*********************************************Fin Solicitud**********************************************")
		} else {
			fmt.Println("Persona encontrada:")
			fmt.Println(*personaBucada)
			fmt.Println("*********************************************Fin Solicitud**********************************************")
		}
	}
}

func crearPersona(nombre *string, edad *int, altura *float64, peso *float64) *Persona {

	if *edad < 0 || *altura < 0 || *peso < 0{
		fmt.Println("Error en alguno de los datos ingresados"); 
		return nil
	}
	return &Persona{Nombre: *nombre, Edad: *edad, Altura: *altura, Peso: *peso}

}

func ordenarPersonas(personas []Persona, criterio int) []Persona{
	switch criterio{
		case 1:
			sort.Slice(personas, func(i, j int) bool {
				return personas[i].Nombre < personas[j].Nombre
			})
		case 2:
			 sort.Slice(personas, func(i, j int) bool {
				return personas[i].Edad < personas[j].Edad
			})
		case 3:
			 sort.Slice(personas, func(i, j int) bool {
				return personas[i].Altura < personas[j].Altura
			})
		case 4:
			 sort.Slice(personas, func(i, j int) bool {
				return personas[i].Peso < personas[j].Peso
			})
		default:
	}

	return personas
}




func calcularIMC(persona Persona) string{
	IMC := persona.Peso/(persona.Altura*persona.Altura)
	var categoria string

	if(IMC < 18.5){
		categoria="Bajo peso"
	} else if(IMC <= 24.9){
		categoria="Peso normal"
	} else if(IMC <= 29.9){
		categoria="Sobrepeso"
	} else{
		categoria="Obesidad"
	}

	return fmt.Sprintf("La persona %s, de %d años de edad, tiene peso de %.2fkg y altura %.2fm. Su IMC es de %.2f, categorizado en %s", persona.Nombre, persona.Edad, persona.Peso, persona.Altura, IMC, categoria)
}




func buscarPersona(personas []Persona, criterio int, valor any) *Persona{
	var persona *Persona

	for _,value := range personas{
		tipo := reflect.TypeOf(value)
	
		for j := 0; j < tipo.NumField(); j++ {
			if j+1 == criterio {
				campo := tipo.Field(j)
				campoValor := reflect.ValueOf(value).FieldByIndex(campo.Index)
				
				if campoValor.Interface() == valor {
					persona = &value
                    break
				}
			}
		}
		if persona != nil {
            break
        } 
	}
	
	return persona
}
func main() {
	var nombre string
	var edad int
	var altura float64
	var peso float64
	var accion int
	var personas []Persona
	var criterio int

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

			persona := crearPersona(&nombre, &edad, &altura, &peso)
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
			ordenarPersonas(personas, criterio)
			fmt.Println(personas)
			fmt.Println("*********************************************Fin Solicitud**********************************************")
			accion = 0
		case 3:
			buscarPersonaPorConsola(&criterio, &nombre, &personas, &edad, &altura, &peso)
			accion = 0
		case 4:
			calcularIMCPorConsola(&criterio, &nombre, &personas, &edad, &altura, &peso)
			accion = 0
		case 5:
			fmt.Println("Gracias, vuelva pronto!")
		default:
			accion = 0
		}
	}

}


