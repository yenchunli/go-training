package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("=======================")
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path {}", r.URL.Path)
	fmt.Println("scheme {}", r.URL.Scheme)

	for k, v := range r.Form {
		fmt.Println("key", k)
		fmt.Println("val", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello world")
}

func main() {
	http.HandleFunc("/people", sayHello)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("Listening", err)
	}
}
