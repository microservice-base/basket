package basketbusinessserviceimpl

import (
	"basket/basket/basketbusinessservice"
	bs "basket/basket/basketbusinessservice"
	"fmt"
)

func AddYap(a bs.Basketinterface) string {
	return a.Add()
}

func Insert(n string, c string) {
	lel := basketbusinessservice.Addstruct{Name: n, Color: c}
	result := AddYap(lel)
	fmt.Println(result)
}
