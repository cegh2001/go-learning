package main

import "fmt"

// closure
	func makeCounter() func() int {
    i := 0
    return func() int {
        i++
        return i
    }
	}

func main() {
	//defer
	defer fmt.Println("Esto se ejecutará al final de la función main.")
	
	// Condicionales
	edad := 20

	// if edad >= 18 {
	// 	fmt.Println("Eres mayor de edad.")
	// }else {
	// 	fmt.Println("Eres menor de edad.")
	// } //modo estandar

	// Modo assertive/negative
	if edad < 18 {
		fmt.Println("Eres menor de edad.")
		return
	}

	fmt.Println("Eres mayor de edad.")

	// Bucles

	// Clasico
	for i := 0; i < 5; i++ {
		fmt.Printf("Iteración %d\n", i)
	}

	// While
	num := 0
	for num < 5 {
		fmt.Printf("Número: %d\n", num)
		num++
	}

	// Infinito con break
	n := 0

	for {
		n++

		if n == 5 {
			continue
	}

	fmt.Printf("Contador infinito: %d\n", n)

	if n >= 10 {
		break
		}
	}

	// ranges
	slice := []string{"Go", "Python", "JavaScript"}

	for index, lenguaje := range slice {
		fmt.Printf("Índice: %d, Lenguaje: %s\n", index, lenguaje)
	}

	// switch
	dia := "Lunes"

	switch dia {
	case "Lunes":
		fmt.Println("Hoy es lunes.")
	case "Martes":
		fmt.Println("Hoy es martes.")
	default:
		fmt.Println("No es ni lunes ni martes.")
	}

	// closure
	counter := makeCounter()

	fmt.Println(counter()) // Imprime: 1
	fmt.Println(counter())
	fmt.Println(counter()) // Imprime: 3
	
}