package main

import "fmt"

type Shape interface {
	Area() float64
}

type Circulo struct {
	Radio float64
}

type Cuadrado struct {
	Lado float64
}

type Triangulo struct {
	Base   float64
	Altura float64
}

func (c *Circulo) Area() float64 {
	return 3.1416 * c.Radio * c.Radio
}

func (c *Cuadrado) Area() float64 {
	return c.Lado * c.Lado
}

func (t *Triangulo) Area() float64 {
	return (t.Base * t.Altura) / 2
}

func mostrarArea(key string, shape Shape){
	if shape.Area() > 20 {
		defer func() {
			cadena := recover()
			fmt.Println(cadena)
		}()
		panic(fmt.Sprintf("Error, área del %s es demasiado grande", key))
	} 
	
	fmt.Printf("El área del %s es: %.2f\n", key, shape.Area())

}
func main() {
	var mapShapes = make(map[string]Shape)

	circulo1 := Circulo{Radio: 3}
	cuadrado1 := Cuadrado{Lado: 4}
	triangulo1 := Triangulo{Base: 2, Altura: 3}

	mapShapes["circulo"] = &circulo1
	mapShapes["cuadrado"] = &cuadrado1
	mapShapes["triangulo"] = &triangulo1

	
	for key, shape := range mapShapes {
		mostrarArea(key, shape)	
	}

}