package switchPractice
import "strings"

func ClasificarSangre(tipoSangre string) string{
	var response string;

	switch(strings.ToUpper(tipoSangre)){
	case "O+":
		response = "Grupo sanguíneo O, factor Rh positivo"
	case "O-":
		response = "Grupo sanguíneo O, factor Rh negativo"
	case "A+":
		response = "Grupo sanguíneo A, factor Rh positivo"
	case "A-":
		response = "Grupo sanguíneo A, factor Rh negativo"
	case "B+":
		response = "Grupo sanguíneo B, factor Rh positivo"
	case "B-":
		response = "Grupo sanguíneo B, factor Rh negativo"
	case "AB+":
		response = "Grupo sanguíneo AB, factor Rh positivo"
	case "AB-":
		response = "Grupo sanguíneo AB, factor Rh negativo"
	default: 
		response = "Tipo de sangre inválido"
	}

	return response
}