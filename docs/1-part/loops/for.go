package main

import "fmt"

func main() {	
	for i := 0; i < 5; i++ {
		fmt.Printf("Valor de i: %d\n", i)
	}
	// for i := range [5]int{} {
		// 	fmt.Printf("Valor de i: %d\n", i)
		// } evitar
		
		// for i := range 5 {
		// 	fmt.Printf("Valor de i: %d\n", i)
		// }
}