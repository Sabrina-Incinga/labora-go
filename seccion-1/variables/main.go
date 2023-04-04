package main
import "seccion-1/variables/restaurantes"
import "seccion-1/variables/sistema-solar"
import "seccion-1/variables/equipos"
import "fmt"

func main(){
	restaurantes.Restaurantes()
	fmt.Println("")
	sistemaSolar.SistemaSolar()
	fmt.Println("")
	equipos.Equipos()
}