package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func writeJSON(w http.ResponseWriter, body map[string]string) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(body)
}

func hello(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, map[string]string{"message": "Hello, World!"})
}

func healthy(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, map[string]string{"status": "healthy"})
}

//teste c2
func main() {
	http.HandleFunc("GET /hello", hello)
	http.HandleFunc("GET /healthy", healthy)

	log.Println("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
