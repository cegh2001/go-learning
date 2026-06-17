package main

import "fmt"

func main() {
	var numeros [5]int // Declaración de un array de enteros con capacidad para 5

	// Asignación de valores a los elementos del array
	numeros[0] = 10
	numeros[1] = 20
	numeros[2] = 30
	numeros[3] = 40
	numeros[4] = 50
	fmt.Println(numeros)

	num := [5]int{1, 2, 3, 4, 5} // Declaración e inicialización de un array
	fmt.Println(num)

	for _, valor := range num { // Iteración sobre el array utilizando range
		fmt.Println(valor)
	}
}