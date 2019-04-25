package main

import (
	"encoding/json"
	"fmt"
	"html"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type myData struct {
	Name    string
	Surname string
}

func list(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path+"-"+r.Method))
}

func save(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)

	var data myData
	err := decoder.Decode(&data)
	if err != nil {
		panic(err)
	}

	name := data.Name
	surname := data.Surname

	fmt.Fprintf(w, r.URL.Path)
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path+" -> "+r.Method+" -> "+name+" -> "+surname))
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome Home Page")
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, r.URL.Path)
	fmt.Fprintln(w, "")
	fmt.Fprintf(w, " Not found page :) ")
}

func apiMethods() string {
	return "Hello, '/basket'"
}

func main() {
	apiMethods()

	r := mux.NewRouter()

	r.NotFoundHandler = http.HandlerFunc(notFound) // auto not found page handler func
	r.HandleFunc("/", home)
	r.HandleFunc("/basket/list", list)
	r.HandleFunc("/basket/save", save)
	log.Fatal(http.ListenAndServe(":8002", r))

	// curl --request GET http://localhost:8002/basket/list
	// curl --request GET http://localhost:8002/basket/save
	// curl --request GET http://localhost:8002/
}
