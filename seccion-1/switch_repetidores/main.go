package main

import (
	"fmt"
	"labora-go/seccion-1/switch_repetidores/repetidores"
	"labora-go/seccion-1/switch_repetidores/switch"
	"time"
)

func main() {
	currentTime := time.Now()
	weekday := int(currentTime.Weekday())
	
	switchPractice.SwitchPractice(weekday)
	repetidores.Repetidores()
	fmt.Println(switchPractice.ClasificarSangre("Ab+"))
}