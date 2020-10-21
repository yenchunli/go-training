package main

import (
	"fmt"
	"log"
	"net/http"
)

// Handler HTTP Request
func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi %s, welcome", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", Handler)
	log.Fatal(http.ListenAndServe(":8080", nil))

}
