package main

import "fmt"

func main() {

	text := "Hello, World!"

	numberOfVowels := 0
	for i := 0; i < len(text); i++ {
		if text[i] == 'a' || text[i] == 'e' || text[i] == 'i' || text[i] == 'o' || text[i] == 'u' {
			numberOfVowels++
		}
	}
	fmt.Printf("The number of vowels in the text is: %d\n", numberOfVowels)
}