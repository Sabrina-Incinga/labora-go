package main

import "fmt"

func main(){
	prueba1 := 1030
	prueba2 := 12045
	prueba3 := 176520

	d1,h1,m1,s1 := segundosADateTime(prueba1)
	d2,h2,m2,s2 := segundosADateTime(prueba2)
	d3,h3,m3,s3 := segundosADateTime(prueba3)

	fmt.Printf("%d segundos son %d día/s, %d hora/s, %d minuto/s con %d segundo/s\n", prueba1, d1, h1, m1, s1)
	fmt.Printf("%d segundos son %d día/s, %d hora/s, %d minuto/s con %d segundo/s\n", prueba2, d2, h2, m2, s2)
	fmt.Printf("%d segundos son %d día/s, %d hora/s, %d minuto/s con %d segundo/s\n", prueba3, d3, h3, m3, s3)

}

func segundosADateTime(segundos int) (int, int, int, int){
	var dias, horas, min, seg int =  0, 0, 0, 0

	
	aux:=segundos/60
	aux2:= aux/60

	seg = segundos%60
	min = aux%60
	horas = aux2%24
	dias = aux2/24
	
	return dias, horas, min, seg
}