package main

import (
	api "basket/basket/api/basketApi"

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
	myCustomRouter.HandleFunc("/basket/list", api.BasketList).Methods("GET")
	// myCustomRouter.HandleFunc("/basket/list/save", api.BasketAddItem).Methods("POST")
	myCustomRouter.HandleFunc("/basket/list/{name}/{color}", api.BasketAddItem).Methods("POST")
	myCustomRouter.HandleFunc("/basket/list/{name}/{color}", api.BasketUpdateItem).Methods("PUT")
	myCustomRouter.HandleFunc("/basket/list/{name}", api.BasketDeleteItem).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8002", myCustomRouter))
}

func main() {
	api.InitialMigration()
	handleRequests()
}
