package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from User-Service")
}

func main() {
	file, err := os.OpenFile("/var/log/user.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler).Methods("GET")
	log.SetOutput(file)

	log.Println("Starting server User-services on :8000")

	http.Handle("/", r)
	fmt.Println("Starting server on :8000")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
