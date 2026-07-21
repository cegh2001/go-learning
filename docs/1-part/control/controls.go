package main

import (
    "errors"
    "fmt"
)

func main() {
    if err := run(); err != nil {
        fmt.Println("Error:", err)
    }
}

func run() error {
    edad := 25

    if edad >= 18 {
        fmt.Println("Eres mayor de edad.")
    }

    if edad < 0 {
        return errors.New("edad no puede ser negativa")
    }

    return nil
}

// func main() {
//     defer func() {
//         if r := recover(); r != nil {
//             fmt.Println("Se recuperó del panic:", r)
//         }
//         fmt.Println("Cierre final de main.")
//     }()

//     edad := 25
//     if edad >= 18 {
//         fmt.Println("Eres mayor de edad.")
//     }

//     if edad < 0 {
//         fmt.Println("La edad es un número negativo.")
//         panic("Edad no puede ser negativa.")
//     }
// } mal manejo de control