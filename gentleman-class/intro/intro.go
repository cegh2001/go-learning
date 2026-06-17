package main

import (
	"fmt"
	"strings"
)

func main() {

	// Números
	entero := 42
	flotante := 3.14
	suma := entero + int(flotante)
	fmt.Printf("La suma de %d y %.2f convertido a entero es: %d\n", entero, flotante, suma)

	// Texto
	cadena := "Hola, Go!"
	concatenacion := cadena + " ¿Cómo estás?"
	enMayusculas := strings.ToUpper(cadena)
	fmt.Println("Cadena original:", cadena)
	fmt.Println("Concatenación:", concatenacion)
	fmt.Println("En mayúsculas:", enMayusculas)

	// Booleanos
	verdadero := true
	falso := false
	fmt.Printf("El valor de verdadero es: %t\n", verdadero)
	fmt.Printf("El valor de falso es: %t\n", falso)

	// Listas: Arrays y Slices
	array := [5]int{1, 2, 3, 4, 5}
	slice := []string{"Go", "Python", "Java"}
	slice = append(slice, "C++")
	fmt.Println("Array:", array)
	fmt.Println("Slice:", slice)

	// Mapas
	similarDiccionario := map[string]int{
		"manzana":  1,
		"banana":   2,
		"naranja":  3,
	}
	fmt.Println("Mapa:", similarDiccionario)

	// Structs
	type Persona struct {
		Nombre string
		Edad   int
	}

	persona := Persona{Nombre: "Alice", Edad: 30}
	fmt.Printf("Persona: %s, Edad: %d\n", persona.Nombre, persona.Edad)
}
	