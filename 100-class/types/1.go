package main

import "fmt"

// We define distinct custom types based on float64
type USD float64
type EUR float64

func main() {
    var walletUSD USD = 100.0
    var walletEUR EUR = 50.0

    // What happens if we try to do this?
    // total := walletUSD + walletEUR
    // fmt.Println(total) 
	// This will result in a compilation error because we cannot directly add two different custom types, even though they are both based on float64.

	total := walletUSD + USD(walletEUR * 1.1)
	fmt.Println(total)
}