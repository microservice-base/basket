package custom

import (
	"basket/basket/basketbusinessservice"
	"basket/basket/basketbusinessserviceimpl"
	domain "basket/basket/domain"

	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"

	// _ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB
var err error

//
func InitialMigration() {

	lel := basketbusinessservice.Addstruct{Sayi: 3}
	basketbusinessserviceimpl.AddYap(lel)

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
func BasketList(w http.ResponseWriter, r *http.Request) {
	fmt.Println("List of items")

	db, err := gorm.Open("sqlite3", "./sqlitetest.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect db")
	}
	defer db.Close()

	var users []domain.BasketEntity
	db.Find(&users)
	json.NewEncoder(w).Encode(users)
}

//
func BasketAddItem(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Added Item")

	db, err := gorm.Open("sqlite3", "./sqlitetest.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect db")
	}
	defer db.Close()

	vars := mux.Vars(r)
	name := vars["name"]
	color := vars["color"]

	db.Create(&domain.BasketEntity{Name: name, Color: color})

	fmt.Fprintf(w, "New user successfully Created")
}

//
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Item deleted")

	db, err := gorm.Open("sqlite3", "./sqlitetest.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect db")
	}
	defer db.Close()

	vars := mux.Vars(r)
	name := vars["name"]

	var users domain.BasketEntity
	db.Where("name = ? ", name).Find(&users)
	db.Delete(&users)

	fmt.Fprintf(w, "New user successfully deleted")
}

//
func BasketUpdateItem(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Item updated")

	db, err := gorm.Open("sqlite3", "./sqlitetest.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect db")
	}
	defer db.Close()

	vars := mux.Vars(r)
	paramName := vars["name"]
	paramColor := vars["color"]

	var users domain.BasketEntity
	db.Where("name = ? ", paramName).Find(&users)

	users.Name = paramName
	users.Color = paramColor
	db.Model(&users).Updates(users)
}
