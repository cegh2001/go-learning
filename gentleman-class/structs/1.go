package main

import "fmt"

type Rectangulo struct {
	Ancho, Alto   float64
}

func (r Rectangulo) Area() float64 {
	return r.Ancho * r.Alto
}

func main() {
	rectangulo := Rectangulo{Ancho: 5.0, Alto: 3.0}
	fmt.Printf("Área del rectángulo: %f\n", rectangulo.Area())
}