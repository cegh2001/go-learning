package main

import "fmt"

func main() {
	miMapa := map[string]int{
		"manzana":  1,
		"banana":   8,
		"naranja":  2,
	}
	fmt.Printf("mi mapa: %v\n", miMapa)

	otroMapa := make(map[string]string)

	otroMapa["nombre"] = "Juan"
	otroMapa["apellido"] = "Pérez"
	fmt.Printf("otro mapa: %v\n", otroMapa)

}