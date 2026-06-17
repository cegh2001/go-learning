package main

import (
	"errors"
	"fmt"
)

// Función clasica
func sumar(a, b int) int {
	return a + b
}

// Función retornando múltiples valores

func dividir(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("No se puede dividir por cero")
	}
	return a / b, nil
}

func main() {
	// Uso de la función sumar
	resultado := sumar(5, 3)
	fmt.Printf("La suma de 5 y 3 es: %d\n", resultado)

	// Uso de la función dividir
	cociente, err := dividir(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Printf("La división de 10 entre 2 es: %.2f\n", cociente)

	// Intentando dividir por cero
	_, err = dividir(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	}
}