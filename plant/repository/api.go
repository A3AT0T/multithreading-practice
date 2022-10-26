package repository

import "github.com/SchoolGolang/multithreading-practice/plant"

type Repository interface {
	AddPlant(data plant.PlantData)
	RemovePlant(id string)
	GetPlant(id string) plant.PlantData
	GetHydration(id string) float64
	GetPh(id string) int
	GetHealth(id string) plant.HealthData
	GetNormalHydration(id string) float64
	GetNormalPh(id string) (int, int)
	GetPlantIds() []string
	GetAge(id string) int
	SetPh(id string, ph int)
	SetHydration(id string, hydration float64)
	SetHealth(id string, data plant.HealthData)

}
