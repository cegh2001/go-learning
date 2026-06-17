package main

import "fmt"

func main() {
	listNumbers := []int{1, 2, 3, 4, 5, 11, 12, 14}

	saveNumbers := func(numbers []int) {
		fmt.Println("Saving numbers:", numbers)
	}

	// for i := 0; i < len(listNumbers); i++ {
	// 	if i%2 == 0 && i > 10 {
	// 		saveNumbers(listNumbers[:i+1])
	// 	}
	// } error porque recorro indice no valor

	filteredNumbers := []int{}

	for _, number := range listNumbers {
		if number%2 == 0 && number > 10 {
			filteredNumbers = append(filteredNumbers, number)
		}
	}
	saveNumbers(filteredNumbers)
}