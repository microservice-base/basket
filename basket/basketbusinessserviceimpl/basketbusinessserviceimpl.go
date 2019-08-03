package basketbusinessserviceimpl

import (
	"basket/basket/basketbusinessservice"
	"basket/basket/basketdao"
	"basket/basket/basketentity"
)

// sonra yapılacak
func AddYap(a basketbusinessservice.Basketinterface) string {
	return a.Add()
}

//
func InitialMigration() {
	basketdao.InitialMigration()
}

// BasketList return all basket items
func BasketList() []basketentity.BasketEntity {
	return basketdao.BasketList()
}

// BasketAddItem add new item
func BasketAddItem(name *string, color *string) {
	basketdao.BasketAddItem(name, color)
}

// BasketDeleteItem delete item
func BasketDeleteItem(name *string) {
	basketdao.BasketDeleteItem(name)
}

// BasketUpdateItem update item
func BasketUpdateItem(name *string, color *string) {
	basketdao.BasketUpdateItem(name, color)
}
