package repetidores

import ("fmt"
		"strings"
		"github.com/mozillazg/go-unidecode")

func Repetidores() {
	planetas := [9]Planeta{{nombre: "Mercurio", numeroDeLunas: 0},
							{nombre: "Venus", numeroDeLunas: 0},
							{nombre: "Tierra", numeroDeLunas: 1},
							{nombre: "Marte", numeroDeLunas: 2},
							{nombre: "Júpiter", numeroDeLunas: 79},
							{nombre: "Saturno", numeroDeLunas: 82},
							{nombre: "Urano", numeroDeLunas: 27},
							{nombre: "Neptuno", numeroDeLunas: 14},
							{nombre: "Plutón", numeroDeLunas: 5}}

	var nombrePlaneta string
	fmt.Println("Ingrese el nombre del planeta cuyo número de lunas desee saber: ")
	fmt.Scanln(&nombrePlaneta)
	fmt.Println(mostrarLunas(planetas[:], nombrePlaneta))
}

type Planeta struct {
	nombre        string
	numeroDeLunas int
}

func mostrarLunas(planetas []Planeta, nombrePlaneta string) string{
	var respuesta string
	
	for i := 0; i < len(planetas); i++ {
		var plnt Planeta = planetas[i]

		if strings.EqualFold(unidecode.Unidecode(plnt.nombre),unidecode.Unidecode(nombrePlaneta)) {
			respuesta = fmt.Sprintf("El número de lunas del planeta %s es: %d", plnt.nombre, plnt.numeroDeLunas)
			break
		} else { respuesta = "Planeta no encontrado"}
	}

	return respuesta
}
