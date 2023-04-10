package main

import (
	"labora-go/seccion-1/switch_repetidores/repetidores"
	"labora-go/seccion-1/switch_repetidores/switch"
	"fmt"
)

func main() {
	switchPractice.SwitchPractice()
	repetidores.Repetidores()
	fmt.Println(switchPractice.ClasificarSangre("Ab+"))
}