package basketbusinessservice

type Basketinterface interface {
	 Add() string
	// list(value string) int
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
