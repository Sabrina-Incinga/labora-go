package peliculas

import "fmt"

func Peliculas(){
	peliculas := [10] string {"Lo que el viento se llevó", "Rápidos y furiosos"}

	fmt.Println(peliculas)
	fmt.Println(peliculas[1])
	fmt.Print(len(peliculas))
}