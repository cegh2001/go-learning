package main

import "fmt"

func main() {
	const pi float64 = 3.14159
	fmt.Printf("The value of pi is: %f\n", pi)

	var (
		e float64 = 2.71828
		g float64 = 9.81
	)
	fmt.Printf("The value of e is: %f\n", e)
	fmt.Printf("The value of g is: %f\n", g)

	g = 10.0
	fmt.Printf("The new value of g is: %f\n", g)
}