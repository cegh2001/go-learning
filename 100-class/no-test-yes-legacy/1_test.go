// math_test.go
package main

import "testing"

// Test functions must always start with the word "Test"
func TestAdd(t *testing.T) {
    result := Add(2, 3)
    expected := 5

    if result != expected {
        t.Errorf("Result was %d, but expected %d", result, expected)
    }
}