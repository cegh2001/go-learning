//Over-engineered ("Trying too hard to be clean")
// Creating interfaces, factories, and multiple files for a simple task:

// package main

// import "fmt"

// type StringPrinter interface {
//     Print(s string)
// }

// type ConsolePrinter struct{}

// func (cp ConsolePrinter) Print(s string) {
//     fmt.Println(s)
// }

// type PrinterFactory struct{}
// ... 20 more lines just to print a greeting!

// Pragmatic Go

package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}
