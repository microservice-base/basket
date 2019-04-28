package custom

import (
	"basket/basket/basketbusinessserviceimpl"
	"basket/basket/dao/basketDao"

	"net/http"
)

//
func InitialMigration() {

	basketbusinessserviceimpl.Insert("3", "4")

	basketDao.InitialMigration()
}

//
func BasketList(w http.ResponseWriter, r *http.Request) {
	basketDao.BasketList(w, r)
}

//
func BasketAddItem(w http.ResponseWriter, r *http.Request) {
	basketDao.BasketAddItem(w, r)
}

//
func BasketDeleteItem(w http.ResponseWriter, r *http.Request) {
	basketDao.BasketDeleteItem(w, r)
}

//
func BasketUpdateItem(w http.ResponseWriter, r *http.Request) {
	basketDao.BasketUpdateItem(w, r)
}
