package custom

import (
	"basket/basket/basketbusinessservice"
	"basket/basket/basketbusinessserviceimpl"
	"basket/basket/dao/basketDao"
	"fmt"

	"net/http"
)

//
func InitialMigration() {
	basketDao.InitialMigration()
}

//
func BasketList(w http.ResponseWriter, r *http.Request) {
	basketDao.BasketList(w, r)
}

//
func BasketAddItem(w http.ResponseWriter, r *http.Request) {
	n := "2"
	c := "3"

	lel := basketbusinessservice.Addstruct{Name: n, Color: c}
	result := basketbusinessserviceimpl.AddYap(lel)
	fmt.Println(result)

	basketbusinessserviceimpl.BasketAddItem(w, r)

}

//
func BasketDeleteItem(w http.ResponseWriter, r *http.Request) {
	basketDao.BasketDeleteItem(w, r)
}

//
func BasketUpdateItem(w http.ResponseWriter, r *http.Request) {
	basketDao.BasketUpdateItem(w, r)
}
