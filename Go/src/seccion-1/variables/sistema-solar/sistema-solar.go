package sistemaSolar
import "fmt"


func SistemaSolar(){

	var (planetName string = "Mercurio"
        planetName2 string = "Venus"
        planetName3 string = "Tierra"
        planetName4 string = "Marte"
        planetName5 string = "Júpiter"
        planetName6 string = "Saturno"
        planetName7 string = "Urano"
        planetName8 string = "Neptuno"
        planetName9 string = "Plutón")
	var numberOfMoons, numberOfMoons2, numberOfMoons3, numberOfMoons4, numberOfMoons5, numberOfMoons6, numberOfMoons7, numberOfMoons8 = 0, 1, 2, 79, 82, 27, 14, 5

	fmt.Printf("%s tiene %d lunas\n", planetName, numberOfMoons)
	fmt.Printf("%s tiene %d lunas\n", planetName2, numberOfMoons)
	fmt.Printf("%s tiene %d luna\n", planetName3, numberOfMoons2)
	fmt.Printf("%s tiene %d lunas\n", planetName4, numberOfMoons3)
	fmt.Printf("%s tiene %d lunas\n", planetName5, numberOfMoons4)
	fmt.Printf("%s tiene %d lunas\n", planetName6, numberOfMoons5)
	fmt.Printf("%s tiene %d lunas\n", planetName7, numberOfMoons6)
	fmt.Printf("%s tiene %d lunas\n", planetName8, numberOfMoons7)
	fmt.Printf("%s tiene %d lunas\n", planetName9, numberOfMoons8)



}