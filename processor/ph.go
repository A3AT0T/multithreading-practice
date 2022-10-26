package processor

import (
	"context"
	droneRepository "github.com/SchoolGolang/multithreading-practice/drone/repository"
	"github.com/SchoolGolang/multithreading-practice/plant/repository"
	"github.com/SchoolGolang/multithreading-practice/sensor"
	"log"
)

type PHProcessor struct {
	plantsRepo repository.Repository
	input      <-chan sensor.SensorData[int]
	dronesRepo droneRepository.DroneRepo
}

func NewPHProcessor(
	plantsRepo repository.Repository,
	input <-chan sensor.SensorData[int],
	dronesRepo droneRepository.DroneRepo,
) *PHProcessor {
	return &PHProcessor{
		plantsRepo: plantsRepo,
		input:      input,
		dronesRepo: dronesRepo,
	}
}

func (p *PHProcessor) RunProcessor(ctx context.Context) {
	for {
		select {
		case recData := <-p.input:

			plant := p.plantsRepo.GetPlant(recData.PlantID)

			switch {
			case recData.Data < plant.NormalLowerPh || recData.Data > plant.NormalUpperPh:
				p.dronesRepo.AdjustSoils(recData.PlantID, (plant.NormalUpperPh+plant.NormalLowerPh)/2)
				log.Printf("У рослини %s з ID: %s: Стан кислотності: %v, при нормі: %v-%v  Встановити: %v  ", plant.Name, recData.PlantID, recData.Data, plant.NormalLowerPh, plant.NormalUpperPh, (plant.NormalUpperPh+plant.NormalLowerPh)/2)
			default:
				log.Printf("У рослини %s з ID: %s: Стан кислотності: %v, при нормі %v-%v", plant.Name, recData.PlantID, recData.Data, plant.NormalLowerPh, plant.NormalUpperPh)
			}
		case <-ctx.Done():
			return
		}
	}
}
