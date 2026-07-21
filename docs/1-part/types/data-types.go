package main

import "fmt"

func main() {
	var name string = "Alice"
	var age int = 30
	var height float64 = 5.6
	var isStudent bool = false
	var grade rune = 'A'
	var hobbies []string = []string{"Reading", "Traveling", "Cooking"}
	var scores map[string]int = map[string]int{"Math": 90, "Science": 85, "History": 88}
	var data interface{} = "This can hold any type of data"
	var ptr *int = &age
	var structData struct {
		Name string
		Age  int
	}
	structData.Name = "Bob"
	structData.Age = 25
	fmt.Printf("Name: %s\n", name)
	fmt.Printf("Age: %d\n", age)
	fmt.Printf("Height: %.1f\n", height)
	fmt.Printf("Is Student: %t\n", isStudent)
	fmt.Printf("Grade: %c\n", grade)
	fmt.Printf("Hobbies: %v\n", hobbies)
	fmt.Printf("Scores: %v\n", scores)
	fmt.Printf("Data: %v\n", data)
	fmt.Printf("Pointer to Age: %d\n", *ptr)
	fmt.Printf("Struct Data: Name=%s, Age=%d\n", structData.Name, structData.Age)
}
