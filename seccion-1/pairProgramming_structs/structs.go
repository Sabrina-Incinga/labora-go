package main

import "fmt"

func main(){
	persona1 := Persona{
		nombre: "Maria",
		edad: 25,
		ciudad: "Springfield",
		telefono: 11584358,
	}
	persona2 := Persona{
		nombre: "Pepe",
		edad: 37,
		ciudad: "Toronto",
		telefono: 1159879468,
	}
	persona3 := Persona{
		nombre: "Alicia",
		edad: 20,
		ciudad: "Tucumán",
		telefono: 3546981,
	}
	persona4 := Persona{
		nombre: "Juan",
		edad: 40,
		ciudad: "Bogota",
		telefono: 1159879468,
	}
	persona5 := Persona{
		nombre: "Lucía",
		edad: 27,
		ciudad: "Francia",
		telefono: 1654897,
	}

	personas := [] Persona{persona1, persona2, persona3, persona4, persona5}

	fmt.Println(persona1)
	persona1.incrementarEdad()
	fmt.Println(persona1)
	fmt.Println(buscarEdad(personas, 20))
	fmt.Println(buscarEdad(personas, 1))

	personaNueva := crearPersona("Jack", 37, "Toronto", 1354687)
	fmt.Println(personaNueva)
	personaNueva.actualizarTelefono(548987654)
	fmt.Println(personaNueva)
}

type Persona struct{
	nombre string
	edad int
	ciudad string
	telefono int64
}

func (p *Persona) incrementarEdad(){
	p.edad++
}

func buscarEdad(personas []Persona, edad int) *Persona{
	for i := 0; i < len(personas); i++ {
		if personas[i].edad == edad{
			return &personas[i]
		}
	}
	return nil
}

func crearPersona(nombre string, edad int, ciudad string, telefono int64) *Persona{
	var personaNueva Persona = Persona{nombre: nombre, edad: edad, ciudad: ciudad, telefono: telefono}
	puntero := &personaNueva

	return puntero
}

func (p *Persona) actualizarTelefono(telefono int64){
	p.telefono = telefono
}

