package main

import (
    "fmt"
    "net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Hello! This is a stable, boring server.")
}

func main() {
    http.HandleFunc("/", helloHandler)
    // Starting the server on port 8080
    http.ListenAndServe(":8080", nil)
}