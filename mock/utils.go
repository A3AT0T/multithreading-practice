package mock

import (
	"github.com/SchoolGolang/multithreading-practice/plant"
	"github.com/google/uuid"
	"math/rand"
)

func GetPlantData() plant.PlantData {
	plantNames := []string{"Grape", "Tomato", "Cucumber", "Microgreen"}
	hydration := GetHydrationData()
	ph := GetPHData()
	health := GetHealthData()
	age := GetAgeData()

	return plant.PlantData{
		ID:               uuid.New().String(),
		Name:             plantNames[rand.Intn(4)],
		NormalHydration:  hydration,
		NormalUpperPh:    ph + rand.Intn(10),
		NormalLowerPh:    ph - rand.Intn(10),
		CurrentHydration: hydration,
		CurrentPh:        ph,
		CurrentHealth:    health,
		Age:              age,
	}
}

func GetHydrationData() float64 {
	return float64(rand.Intn(90)+10) / 100
}

func GetPHData() int {
	return rand.Intn(40) + 10
}

func GetHealthData() plant.HealthData {
	return plant.HealthData{
		LeavesState: float64(rand.Intn(90)) + 10.0,
		RootsState:  float64(rand.Intn(90)) + 10.0,
	}
}
func GetAgeData() int  {
	return  rand.Intn(20)+10
}

//func GetDroneData() drone.Drone {
//	return drone.NewC3PODrone(
//		uuid.New().String(),
//		100,
//		util.Point{
//			X: float64(util.GetRandomIndex(1000)),
//			Y: float64(util.GetRandomIndex(1000)),
//		},
//	)
//}

