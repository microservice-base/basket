// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"net/http"

// 	"github.com/gorilla/mux"
// 	"github.com/jinzhu/gorm"

// 	// _ "github.com/jinzhu/gorm/dialects/postgres"
// 	_ "github.com/jinzhu/gorm/dialects/sqlite"
// )

// var db *gorm.DB
// var err error

// //
// type Users struct {
// 	gorm.Model
// 	Name  string
// 	Email string
// }

// //
// func InitialMigration() {
// 	fmt.Println("initilize db")
// 	db, err := gorm.Open("sqlite3", "./sqlitetest.db")
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		panic("failed to connect db")
// 	}
// 	defer db.Close()

// 	db.AutoMigrate(&Users{})
// }

// //
// func AllUsers(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("all user")

// 	db, err := gorm.Open("sqlite3", "./sqlitetest.db")
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		panic("failed to connect db")
// 	}
// 	defer db.Close()

// 	var users []Users
// 	db.Find(&users)
// 	json.NewEncoder(w).Encode(users)
// }

// //
// func NewUser(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprint(w, "new user")

// 	db, err := gorm.Open("sqlite3", "./sqlitetest.db")
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		panic("failed to connect db")
// 	}
// 	defer db.Close()

// 	vars := mux.Vars(r)
// 	name := vars["name"]
// 	email := vars["email"]

// 	db.Create(&Users{Name: name, Email: email})

// 	fmt.Fprintf(w, "New user successfully Created")
// }

// //
// func DeleteUser(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprint(w, "delete user")

// 	db, err := gorm.Open("sqlite3", "./sqlitetest.db")
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		panic("failed to connect db")
// 	}
// 	defer db.Close()

// 	vars := mux.Vars(r)
// 	name := vars["name"]

// 	var users Users
// 	db.Where("name = ? ", name).Find(&users)
// 	db.Delete(&users)

// 	fmt.Fprintf(w, "New user successfully deleted")
// }

// //
// func UpdateUser(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprint(w, "update user")

// 	db, err := gorm.Open("sqlite3", "./sqlitetest.db")
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		panic("failed to connect db")
// 	}
// 	defer db.Close()

// 	vars := mux.Vars(r)
// 	paramName := vars["name"]
// 	paramEmail := vars["email"]

// 	var users Users
// 	db.Where("name = ? ", paramName).Find(&users)

// 	users.Name = paramName
// 	users.Email = paramEmail
// 	db.Model(&users).Updates(users)
// }
