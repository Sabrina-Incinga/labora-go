package pp4

import (
	"fmt"
	"labora-go/seccion-2/ejercicio-integrador-1/pp2"
)

func CalcularPesoProm(personas *[]pp2.Persona) {
	var PesoProm, sum float64

	for _, persona := range *personas {
		sum += persona.Peso
	}
	PesoProm = sum / float64(len(*personas))
	fmt.Printf("El peso promedio es: %v \n", PesoProm)
}

func MinyMax(personas []pp2.Persona) {
	pp2.OrdenarPersonas(personas, 2)
	fmt.Printf("La persona %s es la menor con %d años \n", personas[0].Nombre, personas[0].Edad)
	fmt.Printf("La persona %s es la mayor con %d años \n", personas[len(personas)-1].Nombre, personas[len(personas)-1].Edad)
}