package basketbusinessservice

type Basketinterface interface {
	// list(value string) int
	Add() string
	// AddTypeEntity() domain.BasketEntity
	// delete(value string) int
	// update(value string) int
}

//
type Addstruct struct {
	Name  string
	Color string
}

//
func (r Addstruct) Add() string {
	return r.Name + "-" + r.Color
}
