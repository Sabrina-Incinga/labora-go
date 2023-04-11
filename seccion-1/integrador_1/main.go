package main

import ( "fmt"
		"sort")

type Persona struct {
	nombre string
	edad   int
	altura float64
	peso   float64
}

func crearPersona(nombre *string, edad *int, altura *float64, peso *float64) *Persona {

	if *edad < 0 || *altura < 0 || *peso < 0{
		fmt.Println("Error en alguno de los datos ingresados"); 
		return nil
	}
	return &Persona{nombre: *nombre, edad: *edad, altura: *altura, peso: *peso}

}

func ordenarPersonas(personas []Persona, criterio int) []Persona{
	switch criterio{
		case 1:
			sort.Slice(personas, func(i, j int) bool {
				return personas[i].nombre < personas[j].nombre
			})
		case 2:
			 sort.Slice(personas, func(i, j int) bool {
				return personas[i].edad < personas[j].edad
			})
		case 3:
			 sort.Slice(personas, func(i, j int) bool {
				return personas[i].altura < personas[j].altura
			})
		case 4:
			 sort.Slice(personas, func(i, j int) bool {
				return personas[i].peso < personas[j].peso
			})
		default:
	}

	return personas
}

func buscarPersona(personas []Persona, criterio int, valor any) *Persona{
	var persona *Persona

	porNombre := func(personas []Persona, valor any) *Persona{
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
	

	return persona
}

func calcularIMC(persona Persona) string{
	IMC := persona.peso/(persona.altura*persona.altura)
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

	return fmt.Sprintf("La persona %s, de %d años de edad, tiene peso de %.2fkg y altura %.2fm. Su IMC es de %.2f, categorizado en %s", persona.nombre, persona.edad, persona.peso, persona.altura, IMC, categoria)
}

func main() {

	persona1 := Persona{
		nombre: "Maria",
		edad: 25,
		altura: 1.50,
		peso: 50,
	}
	persona2 := Persona{
		nombre: "Pepe",
		edad: 37,
		altura: 1.75,
		peso: 80,
	}
	persona3 := Persona{
		nombre: "Alicia",
		edad: 20,
		altura: 1.75,
		peso: 65,
	}
	persona4 := Persona{
		nombre: "Juan",
		edad: 40,
		altura: 1.85,
		peso: 90,
	}
	persona5 := Persona{
		nombre: "Lucía",
		edad: 27,
		altura: 1.63,
		peso: 60,
	}

	personas := [] Persona{persona1, persona2, persona3, persona4, persona5}

	fmt.Println(ordenarPersonas(personas, 2))
	fmt.Println(buscarPersona(personas, 1, "Juan carlos"))
	fmt.Println(calcularIMC(persona1))
	var algo string
	fmt.Scanln(&algo)
	fmt.Println(algo)

	
}
