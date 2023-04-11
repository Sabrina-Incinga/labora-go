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

func buscarPersona(personas []Persona, criterio int, valor any) *Persona{
	var persona *Persona
/* 	porNombre := func(personas []Persona, valor any) *Persona{
		var persona1 *Persona
		for i := 0; i < len(personas); i++ {
			if personas[i].nombre == valor {
				persona1 = &personas[i]
			} 
		}
		return persona1
	}
	porPeso := func(personas []Persona, valor any) *Persona{
		var persona1 *Persona
		for i := 0; i < len(personas); i++ {
			if personas[i].peso == valor {
				persona1 = &personas[i]
			} 
		}
		return persona1
	}
	porEdad := func(personas []Persona, valor any) *Persona{
		var persona1 *Persona
		for i := 0; i < len(personas); i++ {
			if personas[i].edad == valor {
				persona1 = &personas[i]
			} 
		}
		return persona1
	}
	porAltura := func(personas []Persona, valor any) *Persona{
		var persona1 *Persona
		for i := 0; i < len(personas); i++ {
			if personas[i].altura == valor {
				persona1 = &personas[i]
			} 
		}
		return persona1
	}
	switch criterio{
		case 1:
			persona = porNombre(personas, valor)
		case 2:
			persona = porEdad(personas, valor)
		case 3:
			persona = porAltura(personas, valor)
		case 4:
			persona = porPeso(personas, valor)
	}
	 */

	for _,value := range personas{
		tipo := reflect.TypeOf(value)
		//reflect.TypeOf(value) devuelve el tipo de la variable value en tiempo de ejecución, representado como un objeto reflect.Type; en este caso, representa un tipo Persona
		fmt.Println(reflect.TypeOf(value))
		for j := 0; j < tipo.NumField(); j++ {
			//recorro los campos de tipo, que se corresponden con los campos de Persona, y selecciono aquel cuyo índice coincide con el criterio +1
			if j+1 == criterio {
				campo := tipo.Field(j)
				campoValor := reflect.ValueOf(value).FieldByIndex(campo.Index)
				//reflect.ValueOf(value) devuelve un valor de tipo reflect.value que REPRESENTA el valor de la variable value
				//reflect.ValueOf(value).FieldByIndex(campo.Index) devuelve un valor de tipo reflect.value que REPRESENTA el valor de la variable value en el campo cuyo índice corresponde a por campo.Index
				//reflect.ValueOf(value).FieldByIndex(campo.Index).Interface() devuelve el valor real del campo
				if campoValor.Interface() == valor {
					//comparo el valor del campo con el valor recibido por parámetro
					persona = &value
                    break
				}
			}
		}
		if persona != nil {
            break
        } //Debo agregar una validación dado que, de otra manera, 
		//range se recorre completo sobreescribiendo el valor de persona dado 
		//que no hay un return dentro del for que recorre a tipo y necesito tener return fuera del range para cumplir con la firma de la función
	}
	
	return persona
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

func main() {

	persona1 := Persona{
		Nombre: "Maria",
		Edad: 25,
		Altura: 1.50,
		Peso: 50,
	}
	persona2 := Persona{
		Nombre: "Pepe",
		Edad: 37,
		Altura: 1.75,
		Peso: 80,
	}
	persona3 := Persona{
		Nombre: "Alicia",
		Edad: 20,
		Altura: 1.75,
		Peso: 65,
	}
	persona4 := Persona{
		Nombre: "Juan",
		Edad: 40,
		Altura: 1.85,
		Peso: 90,
	}
	persona5 := Persona{
		Nombre: "Lucía",
		Edad: 27,
		Altura: 1.63,
		Peso: 60,
	}

	personas := [] Persona{persona1, persona2, persona3, persona4, persona5}

	fmt.Println(ordenarPersonas(personas, 2))
	fmt.Println(buscarPersona(personas, 1, "Pepe"))
	fmt.Println(calcularIMC(persona1))
	/* var algo string
	fmt.Scanln(&algo)
	fmt.Println(algo) */

}
