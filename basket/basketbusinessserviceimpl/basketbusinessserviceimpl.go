package basketbusinessserviceimpl

import (
	bs "basket/basket/basketbusinessservice"
	"fmt"
)

func AddYap(a bs.Basketinterface) {
	fmt.Println(a.AddType())
}
