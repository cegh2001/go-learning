package main

import "fmt"

func main() {

	edad := 18

	if edad < 18 {
		fmt.Println("Eres menor de edad")
	} else if edad == 18 {
		fmt.Println("Tienes 18 años")
	} else {
		fmt.Println("Eres mayor de edad")
	}

	// switch

	age := 18

	// switch {
	// 	case age ==18:
	// 		fmt.Println("Tienes 18 años")
	// 	case age < 18:
	// 		fmt.Println("Eres menor de edad")
	// 	case age > 18:
	// 		fmt.Println("Eres mayor de edad")
	// 	default:
	// 		fmt.Println("Edad no válida")
	// } // switch sin valor, para evaluar condiciones booleanas

	switch age {
	case 18:
		fmt.Println("Tienes 18 años")
	case 17:
		fmt.Println("Eres menor de edad")
	case 19:
		fmt.Println("Eres mayor de edad")
	default:
		fmt.Println("Edad no válida")
	} //switch con valor, para comparar con casos específicos
}