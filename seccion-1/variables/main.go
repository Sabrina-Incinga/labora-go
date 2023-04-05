package main
import ("labora-go/seccion-1/variables/restaurantes"
		"labora-go/seccion-1/variables/sistema-solar"
		"labora-go/seccion-1/variables/equipos"
		"fmt")

func main(){
	restaurantes.Restaurantes()
	fmt.Println("")
	sistemaSolar.SistemaSolar()
	fmt.Println("")
	equipos.Equipos()
}