package main

import "fmt"

type User struct {
    Name   string
    Email  string
    Age    int
    Active bool
}

type Greeter interface {
    Greet() string
}

func (u User) Greet() string {
    return fmt.Sprintf("Hello, my name is %s", u.Name)
}

func main() {
    u := User{
        Name:   "Ana",
        Email:  "ana@mail.com",
        Age:    28,
        Active: true,
    }

    fmt.Println(u.Greet())
    
}