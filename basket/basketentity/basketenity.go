package basketentity

import "github.com/jinzhu/gorm"

type BasketEntity struct {
	gorm.Model
	Name  string
	Color string
}
