package drone

import "github.com/SchoolGolang/multithreading-practice/util"

type DroneProcessor interface {
	AdjustSoils(string, int)
	Hydrate(string, float64)
	ReplacePlant(string) string
}

type Drone interface {
	ChangePosition(point util.Point)
	Recharge()

}
