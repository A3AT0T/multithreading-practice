package repository

import (
	"github.com/SchoolGolang/multithreading-practice/drone"
	"github.com/SchoolGolang/multithreading-practice/mock"
	"github.com/SchoolGolang/multithreading-practice/plant/repository"
)

type DroneRepo struct {
	plantsService *mock.PlantsServiceMock
	plantsRepo    repository.Repository

	drones map[string]drone.Drone
}

func NewDroneRepo(plantsService *mock.PlantsServiceMock, plantsRepository repository.Repository) DroneRepo {
	return DroneRepo{
		plantsService: plantsService,
		plantsRepo:    plantsRepository,
		drones:        make(map[string]drone.Drone),
	}
}

func (d *DroneRepo) AdjustSoils(plantId string, ph int) {
	d.plantsService.UpdatePlantPH(plantId, ph)
}

func (d *DroneRepo) Hydrate(plantId string, value float64) {
	d.plantsService.UpdatePlantHydration(plantId, value)
}

func (d *DroneRepo) ReplacePlant(plantId string) string {
	d.plantsService.RemovePlant(plantId)
	return d.plantsService.AddPlant()
}

//Рослина виросла
func (d *DroneRepo) AgeRipePlant(plantId string) string {
	d.plantsService.RemovePlant(plantId)
	return d.plantsService.AddPlant()
}