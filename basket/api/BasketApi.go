package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func productAPIMethod() string {
	http.HandleFunc("/productapi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	return "Hello, '/productapi'"
}

func main() {
	productAPIMethod()
	log.Fatal(http.ListenAndServe(":8002", nil))
}
