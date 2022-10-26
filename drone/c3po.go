package drone

import (
	"fmt"
	"github.com/SchoolGolang/multithreading-practice/util"
)

type C3PODrone struct {
	ID       string
	Charge   float64
	Position util.Point
}

func NewC3PODrone(id string, charge float64, position util.Point) *C3PODrone {
	return &C3PODrone{
		ID:       id,
		Charge:   charge,
		Position: position,
	}
}

//func (d *C3PODrone) ChangePosition(point util.Point) {
//
//}

func (d *C3PODrone) Recharge()  {
 if d.Charge < 20 {
	 fmt.Println("go to charge")
 }
}
