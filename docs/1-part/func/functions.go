package main

import "fmt"

func main() {

	x,y := 10, 20

	resultado := sumar(x,y)

	fmt.Printf("El resultado de sumar %d y %d es: %d\n", x, y, resultado)
}

func sumar(a, b int) int {
	return a + b
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}