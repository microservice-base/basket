package basketdao

import (
	"basket/basket/basketentity"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	// _ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB
var err error

// InitialMigration is initilaze db method
func InitialMigration() {

	fmt.Println("initilize db")
	db, err := gorm.Open("sqlite3", "./sqlitetest.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect db")
	}
	defer db.Close()

	db.AutoMigrate(&basketentity.BasketEntity{})

}

// BasketList is items of basket
func BasketList() []basketentity.BasketEntity {

	db, err := gorm.Open("sqlite3", "./sqlitetest.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect db")
	}
	defer db.Close()

	var users []basketentity.BasketEntity
	db.Find(&users)

	return users
}

// BasketAddItem is add item to basket
func BasketAddItem(name *string, color *string) {

	db, err := gorm.Open("sqlite3", "./sqlitetest.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect db")
	}
	defer db.Close()

	db.Create(&basketentity.BasketEntity{Name: *name, Color: *color})

}

// BasketDeleteItem is delete item to basket
func BasketDeleteItem(name *string) {

	db, err := gorm.Open("sqlite3", "./sqlitetest.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect db")
	}
	defer db.Close()

	var users basketentity.BasketEntity
	db.Where("name = ? ", *name).Find(&users)
	db.Delete(&users)

}

// BasketUpdateItem is update item to basket
func BasketUpdateItem(name *string, color *string) {

	db, err := gorm.Open("sqlite3", "./sqlitetest.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect db")
	}
	defer db.Close()

	var users basketentity.BasketEntity

	db.Where("name = ? ", *name).Find(&users)

	users.Name = *name
	users.Color = *color

	db.Model(&users).Updates(users)
}
