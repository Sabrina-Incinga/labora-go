package mensaje

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

type Mensaje struct{
	Tipo string
	Contenido string
	Procesado chan bool
}

type messageQueue struct {
    emailChan   chan Mensaje
    smsChan     chan Mensaje
    pushChan    chan Mensaje
    emailActivo bool
    smsActivo   bool
    pushActivo  bool
}

func (q *messageQueue) stopEmails() {
    q.emailActivo = false
}

func (q *messageQueue) startEmails() {
    q.emailActivo = true
}

func (q *messageQueue) stopSMS() {
    q.smsActivo = false
}

func (q *messageQueue) startSMS() {
    q.smsActivo = true
}

func (q *messageQueue) stopPush() {
    q.pushActivo = false
}

func (q *messageQueue) startPush() {
    q.pushActivo = true
}

func (q *messageQueue) addMessageToQueue(mensaje Mensaje) {
    switch mensaje.Tipo{
		case "Email":
			q.emailChan <- mensaje
		case "SMS":
			q.smsChan <- mensaje
		case "Push":
			q.pushChan <- mensaje
		default: 
			fmt.Println("Tipo de mensaje no reconocido")
	}
}

func CrearMensaje(tipo string) Mensaje{
    var mensajeContenido string
    fmt.Println("Ingrese el contenido de su mensaje:")
    fmt.Scanln(&mensajeContenido)
    var mensaje Mensaje = Mensaje{Tipo: tipo, Contenido: mensajeContenido}
	mensaje.Procesado = make(chan bool)
	fmt.Println("Procesando mensaje...")

	return mensaje
}
func procesarEmail(queue *messageQueue, mensajes *[]Mensaje, signal chan bool) {
	scanner := bufio.NewScanner(os.Stdin)

	for msg := range queue.emailChan {
		fmt.Println("Presione p para pausar procesamiento de emails o enter para continuar")
		if scanner.Scan() {
			text := scanner.Text()
			if text == "p" {
				queue.stopEmails()
			} else{
				queue.startEmails()
			}
		}

        if queue.emailActivo {
			//Agregar mensaje al historial una vez que fue procesado
			go agregarMensajeAHistorial(mensajes, &msg)

            // Procesar el mensaje
            fmt.Println("Procesando correo:", msg.Contenido)
            time.Sleep(2*time.Second)
            fmt.Println("Correo procesado:", msg.Contenido)

            // Enviar confirmación
            msg.Procesado <- true
			fmt.Println("*****************Procesamiento Emails finalizado*****************")
			signal <- true
        } else {
            // No procesar el mensaje
            fmt.Println("Cola de correos detenida")
			signal <- true
        }
    }
}

func procesarSMS(queue *messageQueue, mensajes *[]Mensaje, signal chan bool){
	scanner := bufio.NewScanner(os.Stdin)

	for msg := range queue.smsChan {
		fmt.Println("Presione p para pausar procesamiento de sms o enter para continuar")
		if scanner.Scan() {
			text := scanner.Text()
			if text == "p" {
				queue.stopSMS()
			} else {
				queue.startSMS()
			}
		}
		
        if queue.smsActivo{
			go agregarMensajeAHistorial(mensajes, &msg)

            fmt.Println("Procesando SMS:", msg.Contenido)
            time.Sleep(2*time.Second)
            fmt.Println("SMS procesado:", msg.Contenido)
			
			msg.Procesado <- true
			fmt.Println("*****************Procesamiento SMS finalizado*****************")
			signal <- true
			} else {
				fmt.Println("Cola de SMS detenida")
				signal <- true
			}
	}
}

func procesarPush(queue *messageQueue, mensajes *[]Mensaje, signal chan bool){
	scanner := bufio.NewScanner(os.Stdin)
	
	for msg := range queue.pushChan {
		fmt.Println("Presione p para pausar procesamiento de mensajes push o enter para continuar")
		if scanner.Scan() {
			text := scanner.Text()
			if text == "p" {
				queue.stopPush()
			} else{
				queue.startPush()
			}
		}
        if queue.pushActivo {
			go agregarMensajeAHistorial(mensajes, &msg)

            fmt.Println("Procesando mensajes push:", msg.Contenido)
			time.Sleep(2*time.Second)
            fmt.Println("Mensaje push procesado:", msg.Contenido)

            msg.Procesado <- true
			fmt.Println("*****************Procesamiento Mensajes push finalizado*****************")
			signal <- true
        } else {
            fmt.Println("Cola de mensajes push detenida")
			signal <- true
        }
    }
}

func agregarMensajeAHistorial(mensajes *[]Mensaje, mensaje *Mensaje) {
	<- mensaje.Procesado
	
	*mensajes = append(*mensajes, *mensaje)

	time.Sleep(time.Second)
	fmt.Println("Historial de mensajes: ", *mensajes)

}


func ProcesarMensaje(){
	var mensajes []Mensaje
	var mensaje Mensaje
	var mensajeTipo int
	signalChannel := make(chan bool)
	queue := messageQueue{emailChan: make(chan Mensaje), smsChan: make(chan Mensaje), pushChan: make(chan Mensaje), emailActivo: true, smsActivo:  true, pushActivo: true}

	for {
		fmt.Println("Ingrese el tipo de mensaje:\n\t1. Email\n\t2. SMS\n\t3. Mensaje push\n\t4. Limpiar historial\n\t5. Salir")
		fmt.Scanln(&mensajeTipo)
		switch mensajeTipo{
			case 1:
				mensaje = CrearMensaje("Email")
				go queue.addMessageToQueue(mensaje)
				go procesarEmail(&queue, &mensajes, signalChannel)
			case 2:
				mensaje = CrearMensaje("SMS")
				go queue.addMessageToQueue(mensaje)
				go procesarSMS(&queue, &mensajes, signalChannel)
			case 3:
				fmt.Println(len(queue.pushChan))
				mensaje = CrearMensaje("Push")
				go queue.addMessageToQueue(mensaje)
				go procesarPush(&queue, &mensajes, signalChannel)
			case 4:
				mensajes = nil
				fmt.Println("Cola de mensajes: ", mensajes)
			case 5:
				fmt.Println("Gracias, vuelva pronto")
				return
			default:
				fmt.Println("Opción incorrecta, intente otra vez")
		}
		<- signalChannel
		fmt.Println("Continuar recibido")
	}

}