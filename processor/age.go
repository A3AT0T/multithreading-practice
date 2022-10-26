package processor

import (
	"context"
	droneRepository "github.com/SchoolGolang/multithreading-practice/drone/repository"
	"github.com/SchoolGolang/multithreading-practice/plant/repository"
	"github.com/SchoolGolang/multithreading-practice/sensor"
	"log"
)

type AgeProcessor struct {
	plantsRepo repository.Repository
	input      <-chan sensor.SensorData[int]
	dronesRepo droneRepository.DroneRepo
}

func NewAgeProcessor(
	plantsRepo repository.Repository,
	input <-chan sensor.SensorData[int],
	dronesRepo droneRepository.DroneRepo,
) *AgeProcessor {
	return &AgeProcessor{
		plantsRepo: plantsRepo,
		input:      input,
		dronesRepo: dronesRepo,
	}
}

func (p *AgeProcessor) RunProcessor(ctx context.Context) {
	for {
		select {
		case recData := <-p.input:
			plant := p.plantsRepo.GetPlant(recData.PlantID)
			switch {
			case recData.Data == plant.Age:
				log.Printf("Рослина %s, з ID: %s: дозріла!!!", plant.Name, recData.PlantID)
			default:
				leftDay := plant.Age - recData.Data // 30 днів термін достигання
				log.Printf("Вік рослини %s, з ID: %s: %v, рослина дозріє через %v днів", plant.Name, recData.PlantID, recData.Data, leftDay)
			}
		case <-ctx.Done():
			return
		}
	}
}
