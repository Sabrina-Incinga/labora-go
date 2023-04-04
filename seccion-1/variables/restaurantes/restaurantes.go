package restaurantes
import "fmt"

func Restaurantes(){
	var restaurante = "Ratatouille"
	var restaurante2 = "Mi restaurante"
	var restaurante3 string 
	restaurante3 = "Patito feo"
	var calificacion, calificacion2, calificacion3 = 2.5, 4.6, 3.4


	fmt.Printf("El restaurante %s tiene una calificación de %v\n",restaurante, calificacion)
	fmt.Printf("El restaurante %s tiene una calificación de %v\n",restaurante2, calificacion2)
	fmt.Printf("El restaurante %s tiene una calificación de %v\n",restaurante3, calificacion3)

}