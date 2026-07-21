package main

import "fmt"

func main() {
	x := 10
	// x = changeValue(x) reasignacion
	changeValue(&x) // paso por referencia
	fmt.Printf("Valor de x: %d", x)
}

func changeValue(x *int) {
	*x = 20
}

// func changeValue(x int) int {
// 	x = 20
// 	return x
// } reasignacion