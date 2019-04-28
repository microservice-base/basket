package basketbusinessserviceimpl

import (
	bs "basket/basket/basketbusinessservice"
	"basket/basket/dao/basketDao"
	"basket/basket/domain"
)

// sonra yapılacak
func AddYap(a bs.Basketinterface) string {
	return a.Add()
}

func InitialMigration() {
	basketDao.InitialMigration()
}

// BasketList return all values
func BasketList() []domain.BasketEntity {
	return basketDao.BasketList()
}

// BasketAddItem add new item
func BasketAddItem(name *string, color *string) {
	basketDao.BasketAddItem(name, color)
}

// BasketDeleteItem delete item
func BasketDeleteItem(name *string) {
	basketDao.BasketDeleteItem(name)
}

// BasketUpdateItem update item
func BasketUpdateItem(name *string, color *string) {
	basketDao.BasketUpdateItem(name, color)
}
