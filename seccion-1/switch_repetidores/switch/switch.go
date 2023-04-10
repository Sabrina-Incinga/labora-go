package switchPractice

import (
	"fmt"
)

func SwitchPractice(numeroDeDia int) {

	fmt.Print("El número corresponde al día: ")

	switch numeroDeDia {
	case 0:
		fmt.Println("Domingo")
	case 1:
		fmt.Println("Lunes")
	case 2:
		fmt.Println("Martes")
	case 3:
		fmt.Println("Miércoles")
	case 4:
		fmt.Println("Jueves")
	case 5:
		fmt.Println("Viernes")
	case 6:
		fmt.Println("Sábado")
	}
}
