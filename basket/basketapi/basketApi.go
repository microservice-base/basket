package basketapi

import (
	"basket/basket/basketbusinessserviceimpl"
	"encoding/json"
	"fmt"

	"net/http"

	"github.com/gorilla/mux"
)

//
func InitialMigration() {
	basketbusinessserviceimpl.InitialMigration()
}

//
func BasketList(w http.ResponseWriter, r *http.Request) {
	result := basketbusinessserviceimpl.BasketList()
	json.NewEncoder(w).Encode(result)
}

//
func BasketAddItem(w http.ResponseWriter, r *http.Request) {
	// interface kullanımına açılacak
	// n := "2"
	// c := "3"
	// lel := basketbusinessservice.Addstruct{Name: n, Color: c}
	// result := basketbusinessserviceimpl.AddYap(lel)
	// fmt.Println(result)

	vars := mux.Vars(r)
	name := vars["name"]
	color := vars["color"]

	basketbusinessserviceimpl.BasketAddItem(&name, &color)

	fmt.Fprintf(w, "New user successfully Created")

}

//
func BasketDeleteItem(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	paramName := vars["name"]

	basketbusinessserviceimpl.BasketDeleteItem(&paramName)

	fmt.Fprintf(w, "New user successfully deleted")
}

//
func BasketUpdateItem(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	paramName := vars["name"]
	paramColor := vars["color"]

	basketbusinessserviceimpl.BasketUpdateItem(&paramName, &paramColor)
}
