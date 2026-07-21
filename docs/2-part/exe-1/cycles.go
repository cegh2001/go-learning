package main

import (
	"fmt"
	"math"
)

// Sqrt calcula la raíz cuadrada de x con una estimación inicial dada (initZ)
// y retorna el resultado y la cantidad de iteraciones que tomó.
func Sqrt(x float64, initZ float64) (float64, int) {
	z := initZ
	for i := 1; ; i++ {
		last := z
		z -= (z*z - x) / (2 * z)
		if math.Abs(z-last) < 1e-15 {
			return z, i
		}
	}
}

func main() {
	testValues := []float64{2, 16, 100, 10000, 1000000}

	for _, x := range testValues {
		fmt.Printf("=== Convergencia para x = %v ===\n", x)

		// Probar con z = 1.0
		_, iter1 := Sqrt(x, 1.0)
		fmt.Printf("Conjetura z = 1.0:   %d iteraciones\n", iter1)

		// Probar con z = x/2
		_, iterHalf := Sqrt(x, x/2.0)
		fmt.Printf("Conjetura z = x/2:   %d iteraciones\n", iterHalf)

		// Probar con z = x
		_, iterX := Sqrt(x, x)
		fmt.Printf("Conjetura z = x:     %d iteraciones\n\n", iterX)
	}
}