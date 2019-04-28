package basketbusinessservice

type Basketinterface interface {
	// list(value string) int
	AddType() int
	// delete(value string) int
	// update(value string) int
}

type Addstruct struct {
	Sayi int
}

func (r Addstruct) AddType() int {
	return r.Sayi
}
