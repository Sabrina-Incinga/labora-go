package pp2

import ( "fmt"
		"sort"
		"reflect")

type Persona struct {
	Nombre string
	Edad   int
	Altura float64
	Peso   float64
}

func CalcularIMCPorConsola(criterio *int, nombre *string, personas *[]Persona, edad *int, altura *float64, peso *float64) {
	fmt.Println("Debe seleccionar una persona para calcular su IMC \nIngrese el número de propiedad por la que desea buscar: \n1. Nombre \n2. Edad \n3. Altura \n4. Peso ")
	fmt.Scanln(criterio)
	switch *criterio {
	case 1:
		fmt.Println("Ingrese el nombre de la persona buscada:")
		fmt.Scanln(nombre)
		personaBucada := BuscarPersona(*personas, *criterio, *nombre)
		if personaBucada == nil {
			fmt.Println("Persona no encontrada")
			fmt.Println("*********************************************Fin Solicitud**********************************************")
		} else {
			fmt.Println("Persona encontrada:")
			fmt.Println(CalcularIMC(*personaBucada))
			fmt.Println("*********************************************Fin Solicitud**********************************************")
		}
	case 2:
		fmt.Println("Ingrese la edad de la persona buscada:")
		fmt.Scanln(edad)
		personaBucada := BuscarPersona(*personas, *criterio, *edad)
		if personaBucada == nil {
			fmt.Println("Persona no encontrada")
			fmt.Println("*********************************************Fin Solicitud**********************************************")
		} else {
			fmt.Println("Persona encontrada:")
			fmt.Println(CalcularIMC(*personaBucada))
			fmt.Println("*********************************************Fin Solicitud**********************************************")
		}
	case 3:
		fmt.Println("Ingrese la altura de la persona buscada:")
		fmt.Scanln(altura)
		personaBucada := BuscarPersona(*personas, *criterio, *altura)
		if personaBucada == nil {
			fmt.Println("Persona no encontrada")
			fmt.Println("*********************************************Fin Solicitud**********************************************")
		} else {
			fmt.Println("Persona encontrada:")
			fmt.Println(CalcularIMC(*personaBucada))
			fmt.Println("*********************************************Fin Solicitud**********************************************")
		}
	case 4:
		fmt.Println("Ingrese el peso de la persona buscada:")
		fmt.Scanln(peso)
		personaBucada := BuscarPersona(*personas, *criterio, *peso)
		if personaBucada == nil {
			fmt.Println("Persona no encontrada")
			fmt.Println("*********************************************Fin Solicitud**********************************************")
		} else {
			fmt.Println("Persona encontrada:")
			fmt.Println(CalcularIMC(*personaBucada))
			fmt.Println("*********************************************Fin Solicitud**********************************************")
		}
	}
}

func BuscarPersonaPorConsola(criterio *int, nombre *string, personas *[]Persona, edad *int, altura *float64, peso *float64) {
	fmt.Println("Ingrese el número de propiedad por la que desea buscar: \n1. Nombre \n2. Edad \n3. Altura \n4. Peso ")
	fmt.Scanln(criterio)
	switch *criterio {
	case 1:
		fmt.Println("Ingrese el nombre de la persona buscada:")
		fmt.Scanln(nombre)
		personaBucada := BuscarPersona(*personas, *criterio, *nombre)
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
		personaBucada := BuscarPersona(*personas, *criterio, *edad)
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
		personaBucada := BuscarPersona(*personas, *criterio, *altura)
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
		personaBucada := BuscarPersona(*personas, *criterio, *peso)
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

func CrearPersona(nombre *string, edad *int, altura *float64, peso *float64) *Persona {

	if *edad < 0 || *altura < 0 || *peso < 0{
		fmt.Println("Error en alguno de los datos ingresados"); 
		return nil
	}
	return &Persona{Nombre: *nombre, Edad: *edad, Altura: *altura, Peso: *peso}

}

func OrdenarPersonas(personas []Persona, criterio int) []Persona{
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




func CalcularIMC(persona Persona) string{
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




func BuscarPersona(personas []Persona, criterio int, valor any) *Persona{
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



