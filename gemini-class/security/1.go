package main

import "fmt"

func checkPassword(password string) (bool, error) {
	if password == "" || password == " " {
		return false, fmt.Errorf("password cannot be empty or just a space")
	}
	if len(password) < 8 {
		return false, fmt.Errorf("password must be at least 8 characters long")
	}
	return true, nil
}

// func main() {
// 	password := "myscrt"
// 	valid, err := checkPassword(password)
// 	if !valid {
// 		fmt.Printf("Password validation failed: %s\n", err)
// 	} else {
// 		fmt.Println("Password is valid.")
// 	}
// }

func main() {
    password := "myscrt"

    if _, err := checkPassword(password); err != nil {
        fmt.Printf("Password validation failed: %v\n", err)
        return
    }

    fmt.Println("Password is valid.")
}