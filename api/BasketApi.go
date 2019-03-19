package main

import (
    "fmt"
    "html"
    "log"
    "net/http"
)

func main() {
	http.HandleFunc("/productapi",
		func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	log.Fatal(http.ListenAndServe(":8002", nil))

}
