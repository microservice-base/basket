package main

import (
	"basket/basket/basketapi"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func applicationInit(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Basket Application Running ....")
}

func handleRequests() {
	myCustomRouter := mux.NewRouter().StrictSlash(true)

	myCustomRouter.HandleFunc("/basket", applicationInit).Methods("GET")
	myCustomRouter.HandleFunc("/basket/list", basketapi.BasketList).Methods("GET")
	// myCustomRouter.HandleFunc("/basket/list/save", api.BasketAddItem).Methods("POST")
	myCustomRouter.HandleFunc("/basket/list/{name}/{color}", basketapi.BasketAddItem).Methods("POST")
	myCustomRouter.HandleFunc("/basket/list/{name}/{color}", basketapi.BasketUpdateItem).Methods("PUT")
	myCustomRouter.HandleFunc("/basket/list/{name}", basketapi.BasketDeleteItem).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8002", myCustomRouter))
}

func main() {
	basketapi.InitialMigration()
	handleRequests()
}
