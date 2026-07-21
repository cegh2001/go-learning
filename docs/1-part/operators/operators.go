package main

import (
	"fmt"
	"math"
)

func main() {

	var a int = 5
	var b float32 = 4.32
	const pi float64 = 3.1415139475
	x, y := 14, 15

	fmt.Printf("valor a + b: %f\n", float32(a)+b)
	fmt.Printf("valor a - b: %f\n", float32(a)-b)
	fmt.Printf("valor a * b: %f\n", float32(a)*b)
	fmt.Printf("valor a / b: %f\n", float32(a)/b)
	fmt.Printf("valor a mod b: %f\n", math.Mod(float64(a), float64(b)))
	fmt.Printf("valor pi: %f\n", pi)
	fmt.Printf("valor x,y: %d, %d", x, y)
}