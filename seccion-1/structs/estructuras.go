package main

import "fmt"

func main(){
	persona1 := Persona{
		nombre: "Maria",
		edad: 25,
		ciudad: "Springfield",
		telefono: "+11584358",
	}
	persona2 := Persona{
		nombre: "Pepe",
		edad: 37,
		ciudad: "Toronto",
		telefono: "+1159879468",
	}

	fmt.Println("Valores previos al cambio de ciudad")
	fmt.Println(persona1)
	fmt.Println(persona2)


	fmt.Println("Datos de persona1 después del cambio")
	persona1.cambiarCiudad("Chicago")
	fmt.Println(persona1)

	fmt.Println("Datos de persona1 después intentar cambiar la ciudad repitiendo ciudad")
	persona1.cambiarCiudad("Chicago")
	fmt.Println(persona1)

}

type Persona struct{
	nombre string
	edad int
	ciudad string
	telefono string
}

func (p *Persona) cambiarCiudad(ciudad string){
	if(p.ciudad != ciudad){
		println("Ciudad cambiada")
		p.ciudad = ciudad
	}
}