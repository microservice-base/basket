package basketDao

import (
	"basket/basket/domain"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	// _ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB
var err error

func InitialMigration() {

	fmt.Println("initilize db")
	db, err := gorm.Open("sqlite3", "./sqlitetest.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect db")
	}
	defer db.Close()

	db.AutoMigrate(&domain.BasketEntity{})
}

//
func BasketList() []domain.BasketEntity {

	db, err := gorm.Open("sqlite3", "./sqlitetest.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect db")
	}
	defer db.Close()

	var users []domain.BasketEntity
	db.Find(&users)

	return users
}

//
func BasketAddItem(name *string, color *string) {

	db, err := gorm.Open("sqlite3", "./sqlitetest.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect db")
	}
	defer db.Close()

	db.Create(&domain.BasketEntity{Name: *name, Color: *color})

}

//
func BasketDeleteItem(name *string) {

	db, err := gorm.Open("sqlite3", "./sqlitetest.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect db")
	}
	defer db.Close()

	var users domain.BasketEntity
	db.Where("name = ? ", *name).Find(&users)
	db.Delete(&users)

}

//
func BasketUpdateItem(name *string, color *string) {

	db, err := gorm.Open("sqlite3", "./sqlitetest.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect db")
	}
	defer db.Close()

	var users domain.BasketEntity

	db.Where("name = ? ", *name).Find(&users)

	users.Name = *name
	users.Color = *color

	db.Model(&users).Updates(users)
}
