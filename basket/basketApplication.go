package main

import (
	apis "basket/basket/apis"

	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func applicationInit(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Basket Application Running ....")
}

func handleRequests() {
	myCustomRouter := mux.NewRouter().StrictSlash(true)

	myCustomRouter.HandleFunc("/", applicationInit).Methods("GET")
	myCustomRouter.HandleFunc("/list", apis.BasketList).Methods("GET")
	myCustomRouter.HandleFunc("/list/{name}/{color}", apis.BasketAddItem).Methods("POST")
	myCustomRouter.HandleFunc("/list/{name}/{color}", apis.BasketUpdateItem).Methods("PUT")
	myCustomRouter.HandleFunc("/list/{name}", apis.DeleteUser).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8002", myCustomRouter))
}

func main() {

	apis.InitialMigration()

	handleRequests()

}
