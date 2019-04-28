package basketbusinessserviceimpl

import (
	bs "basket/basket/basketbusinessservice"
	"basket/basket/dao/basketDao"
	"net/http"
)

func AddYap(a bs.Basketinterface) string {
	return a.Add()
}

func BasketAddItem(w http.ResponseWriter, r *http.Request) {

	basketDao.BasketAddItem(w, r)

}
