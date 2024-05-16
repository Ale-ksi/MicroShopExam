// main.go
package main

import (
    "fmt"
    "net/http"
    "github.com/gorilla/mux"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello from User-Service")
}

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/", homeHandler).Methods("GET")

    http.Handle("/", r)
    fmt.Println("Starting server on :8000")
    if err := http.ListenAndServe(":8000", nil); err != nil {
        fmt.Printf("Error starting server: %s\n", err)
    }
}
