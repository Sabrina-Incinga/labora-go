package mensaje

import (
	"fmt"
	"time"
)

type Mensaje struct{
	Tipo string
	Contenido string
}

func CrearMensaje(tipo string) Mensaje{
    var mensajeContenido string
    fmt.Println("Ingrese el contenido de su mensaje:")
    fmt.Scanln(&mensajeContenido)
    var mensaje Mensaje = Mensaje{tipo, mensajeContenido}
	fmt.Println("Procesando mensaje...")

	return mensaje
}
func procesarEmail(ch chan<- Mensaje, mensajes *[]Mensaje) {
    mensaje := CrearMensaje("Email")
	
	ch <- mensaje
	time.Sleep(time.Second)
    fmt.Println("Email enviado")
}

func procesarSMS(ch chan<- Mensaje, mensajes *[]Mensaje){
	var mensaje Mensaje = CrearMensaje("SMS")
	
	ch <- mensaje
	time.Sleep(time.Second)
	fmt.Println("SMS enviado")
}

func agregarMensajeAHistorial(mensajes *[]Mensaje, mensaje Mensaje) {
	*mensajes = append(*mensajes, mensaje)

	time.Sleep(time.Second)
	fmt.Println("Historial de mensajes: ", *mensajes)
}

func procesarPush(ch chan Mensaje, mensajes *[]Mensaje){
	var mensaje Mensaje = CrearMensaje("Push")
	
	ch <- mensaje
	time.Sleep(time.Second)
	fmt.Println("Mensaje push enviado")
}

func ProcesarMensaje(){
	var mensajes []Mensaje
	var mensajeTipo int
	var mensaje Mensaje
	ch := make(chan Mensaje)

	for {
		fmt.Println("Ingrese el tipo de mensaje:\n\t1. Email\n\t2. SMS\n\t3. Mensaje push\n\t4. Limpiar historial\n\t5. Salir")
		fmt.Scanln(&mensajeTipo)
		switch mensajeTipo{
			case 1:
				go procesarEmail(ch, &mensajes)
				mensaje = <- ch
				go agregarMensajeAHistorial(&mensajes, mensaje)
			case 2:
				go procesarSMS(ch, &mensajes)
				mensaje = <- ch
				go agregarMensajeAHistorial(&mensajes, mensaje)
			case 3:
				go procesarPush(ch, &mensajes)
				mensaje = <- ch
				go agregarMensajeAHistorial(&mensajes, mensaje)
			case 4:
				mensajes = nil
				fmt.Println("Cola de mensajes: ", mensajes)
			case 5:
				fmt.Println("Gracias, vuelva pronto")
				return
			default:
				fmt.Println("Opción incorrecta, intente otra vez")
		}
	}

}