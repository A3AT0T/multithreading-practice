package processor

import (
	"context"
	droneRepository "github.com/SchoolGolang/multithreading-practice/drone/repository"
	"github.com/SchoolGolang/multithreading-practice/plant"
	"github.com/SchoolGolang/multithreading-practice/plant/repository"
	"github.com/SchoolGolang/multithreading-practice/sensor"
	"log"
)

type HealthProcessor struct {
	plantsRepo repository.Repository
	input      <-chan sensor.SensorData[plant.HealthData]
	dronesRepo droneRepository.DroneRepo
}

func NewHealthProcessor(
	plantsRepo repository.Repository,
	input <-chan sensor.SensorData[plant.HealthData],
	dronesRepo droneRepository.DroneRepo,
) *HealthProcessor {
	return &HealthProcessor{
		plantsRepo: plantsRepo,
		input:      input,
		dronesRepo: dronesRepo,
	}
}

func (p *HealthProcessor) RunProcessor(ctx context.Context) {
	for {
		select {
		case recData := <-p.input:

			plant := p.plantsRepo.GetPlant(recData.PlantID)

			switch {
			case plant.CurrentHealth.LeavesState < 50 || plant.CurrentHealth.RootsState < 50:
				p.dronesRepo.ReplacePlant(recData.PlantID)
				log.Printf("Рівень здоров'я рослини %s з ID: %s <50  потребує заміни", plant.Name, recData.PlantID)
			case recData.Data.LeavesState < plant.CurrentHealth.LeavesState:
				p.dronesRepo.ReplacePlant(recData.PlantID)
				log.Printf("Рівень здоров'я листя рослини %s з ID: %s нижче за задовільний,  %f < %f  потребує заміни", plant.Name, recData.PlantID, recData.Data.LeavesState, plant.CurrentHealth.LeavesState)
			case recData.Data.RootsState < plant.CurrentHealth.RootsState:
				p.dronesRepo.ReplacePlant(recData.PlantID)
				log.Printf("Рівень здоров'я коренів рослини %s з ID: %s нижче за задовільний,  %f < %f  потребує заміни", plant.Name, recData.PlantID, recData.Data.RootsState, plant.CurrentHealth.RootsState)
			default:
				log.Printf("У рослини %s з ID: %s рівень здоров'я листя %v у та коріння: %v, ", plant.Name, recData.PlantID, plant.CurrentHealth.LeavesState, plant.CurrentHealth.RootsState)
			}
		case <-ctx.Done():
			return
		}
	}
}
