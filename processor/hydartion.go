package processor

import (
	"context"
	droneRepository "github.com/SchoolGolang/multithreading-practice/drone/repository"
	"github.com/SchoolGolang/multithreading-practice/plant/repository"
	"github.com/SchoolGolang/multithreading-practice/sensor"
	"log"
	//"log"
)

type HydrationProcessor struct {
	plantsRepo repository.Repository
	input      <-chan sensor.SensorData[float64]
	dronesRepo droneRepository.DroneRepo
}

func NewHydrationProcessor(
	plantsRepo repository.Repository,
	input <-chan sensor.SensorData[float64],
	dronesRepo droneRepository.DroneRepo,
) *HydrationProcessor {
	return &HydrationProcessor{
		plantsRepo: plantsRepo,
		input:      input,
		dronesRepo: dronesRepo,
	}
}

func (p *HydrationProcessor) RunProcessor(ctx context.Context) {
	for {
		select {
		case recData := <-p.input:
			plant := p.plantsRepo.GetPlant(recData.PlantID)

			switch {
			case recData.Data < plant.NormalHydration:
				p.dronesRepo.Hydrate(recData.PlantID, plant.NormalHydration)
				log.Printf("У рослини %s з ID: %s: Стан гідрації: %v, при нормі: %v. Встановити норму гідрації: %v ", plant.Name, recData.PlantID, recData.Data, plant.NormalHydration, plant.NormalHydration)
			case float64(plant.NormalHydration) == 0:
				log.Printf("не вірно вказані данні")
			default:
				log.Printf("У рослини %s з ID: %s: стан гідрації: %v, при нормі: %v", plant.Name, recData.PlantID, recData.Data, plant.NormalHydration)
			}
		case <-ctx.Done():
			return
		}
	}

}
