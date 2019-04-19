package main

import (
	"encoding/json"
	"fmt"
	"html"
	"log"
	"net/http"
)

type myData struct {
	Name    string
	Surname string
}

func listHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path+"-"+r.Method))
}

func saveHandler(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)

	var data myData
	err := decoder.Decode(&data)
	if err != nil {
		panic(err)
	}

	name := data.Name
	surname := data.Surname

	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path+" -> "+r.Method+" -> "+name+" -> "+surname))
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Page")
}

func apiMethods() string {
	http.HandleFunc("/", home)
	http.HandleFunc("/basket/list", listHandler)
	http.HandleFunc("/basket/save", saveHandler)
	return "Hello, '/basket'"
}

func main() {
	apiMethods()
	log.Fatal(http.ListenAndServe(":8002", nil))
}
