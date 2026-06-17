package main

import (
	"fmt"
)

// func main() {

// 	fmt.Print("Enter your name: ")
// 	scanner := bufio.NewScanner(os.Stdin)
// 	scanner.Scan()
// 	name := scanner.Text()
// 	fmt.Printf("Hello, %s!\n", name)
// }

func main() {
    var nombre string
    fmt.Print("Tu nombre: ")
    fmt.Scanln(&nombre)
    fmt.Println("Hola,", nombre)
}