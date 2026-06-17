package main

import (
	"fmt"
)

func main() {
	x := 10
	fmt.Printf("x es %d \n", x)
	p := &x
	*p = 21
	fmt.Printf("La dirección de memoria de x es %v", *p)
}