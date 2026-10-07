package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Greeting struct {
	Message string `json:"message"`
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	var message string
	if name == "" {
		message = "Hello, World!"
	} else {
		message = fmt.Sprintf("Hello, %s!", name)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Greeting{Message: message})
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /hello", helloHandler)

	log.Println("greeter listening on :9090")
	log.Fatal(http.ListenAndServe(":9090", mux))
}
