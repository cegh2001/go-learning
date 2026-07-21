package main

import "fmt"

func main() {
    languages := []string{"Go", "Python", "Java", "C++", "JavaScript", "Ruby"}

    // Explicit and crystal clear loop
    for _, lang := range languages {
        fmt.Printf("Im practicing %s\n", lang)
    }
}