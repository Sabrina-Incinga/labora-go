package equipos

import "fmt"

func Equipos() {
	const EQUIPO, EQUIPO2, EQUIPO3, EQUIPO4, EQUIPO5 = "Cebollitas", "Plato Juniors", "Pepe Argento", "Patito Feo FC", "Ovnis FC"
	var tit int = 15
	var tit2, tit3, tit4, tit5 = 35, 23, 12, 17

	fmt.Printf("El equipo %s ha ganado %d títulos\n", EQUIPO, tit)
	fmt.Printf("El equipo %s ha ganado %d títulos\n", EQUIPO2, tit2)
	fmt.Printf("El equipo %s ha ganado %d títulos\n", EQUIPO3, tit3)
	fmt.Printf("El equipo %s ha ganado %d títulos\n", EQUIPO4, tit4)
	fmt.Printf("El equipo %s ha ganado %d títulos\n", EQUIPO5, tit5)

}
