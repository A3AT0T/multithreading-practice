package mock

import (
	"context"
	"github.com/SchoolGolang/multithreading-practice/plant"
	"github.com/SchoolGolang/multithreading-practice/plant/repository"
	"github.com/SchoolGolang/multithreading-practice/sensor"
	sensorRepo "github.com/SchoolGolang/multithreading-practice/sensor/repository"
	"github.com/SchoolGolang/multithreading-practice/util"
	"github.com/google/uuid"
	//"log"

	"math/rand"
	"time"
)

type PlantsServiceMock struct {
	plantsRepo    *repository.PlantsRepository
	phRepo        *sensorRepo.SensorRepo[int]
	hydrationRepo *sensorRepo.SensorRepo[float64]
	healthRepo    *sensorRepo.SensorRepo[plant.HealthData]
	ageRepo		  *sensorRepo.SensorRepo[int]
	frequency int
}

func NewPlantsServiceMock(
	plantsRepo *repository.PlantsRepository,
	phRepo *sensorRepo.SensorRepo[int],
	hydrationRepo *sensorRepo.SensorRepo[float64],
	healthRepo *sensorRepo.SensorRepo[plant.HealthData],
	ageRepo *sensorRepo.SensorRepo[int],
	frequency int,
) *PlantsServiceMock {
	return &PlantsServiceMock{
		plantsRepo:    plantsRepo,
		phRepo:        phRepo,
		hydrationRepo: hydrationRepo,
		healthRepo:    healthRepo,
		ageRepo:	   ageRepo,
		frequency:     frequency,
	}
}

func (ps *PlantsServiceMock) SendRandomUpdates(ctx context.Context) {
	for {
		plantIDs := ps.plantsRepo.GetPlantIds()
		plantID := plantIDs[util.GetRandomIndex(len(plantIDs))]
		select {
		case <-ctx.Done():
			return
		default:
			switch rand.Intn(100) {
			case 0:
				ps.UpdatePlantPH(plantID, GetPHData())
		 	case 1:
				ps.UpdatePlantHydration(plantID, GetHydrationData())
				//log.Printf("plant_id: %s, hd: %v",plantID,GetHydrationData() )
			case 2:
				//log.Printf("plant_id: %s, hp: %v",plantID, GetHealthData() )
				ps.UpdatePlantHealth(plantID, GetHealthData())
			case 3:
				ps.UpdatePlantAge(plantID, GetAgeData() )
				//log.Printf("plant_id: %s, age: %v",plantID, GetAgeData() )
			}
			time.Sleep(10 * time.Millisecond)
		}
	}
}

func (ps *PlantsServiceMock) AddPlant() string {
	plantData := GetPlantData()
	plantId := plantData.ID
	ps.plantsRepo.AddPlant(plantData)

	phSensor := sensor.NewSensor[int](uuid.New().String(), plantId)
	ps.phRepo.AddSensor(phSensor)

	hydrationSensor := sensor.NewSensor[float64](uuid.New().String(), plantId)
	ps.hydrationRepo.AddSensor(hydrationSensor)

	healthSensor := sensor.NewSensor[plant.HealthData](uuid.New().String(), plantId)
	ps.healthRepo.AddSensor(healthSensor)

	ageSensor := sensor.NewSensor[int](uuid.New().String(), plantId) // AGE
	ps.ageRepo.AddSensor(ageSensor)						        	 // AGE

	return plantId
}

func (ps *PlantsServiceMock) RemovePlant(plantID string) {
	ps.plantsRepo.RemovePlant(plantID)

	ps.phRepo.RemoveSensorByPlantID(plantID)
	ps.hydrationRepo.RemoveSensorByPlantID(plantID)
	ps.healthRepo.RemoveSensorByPlantID(plantID)
}

func (ps *PlantsServiceMock) UpdatePlantPH(plantId string, ph int) {
	s := ps.phRepo.GetSensorByPlantID(plantId)
	s.Connect() <- sensor.SensorData[int]{
		PlantID: s.PlantID,
		Data:    ph,
	}
}

func (ps *PlantsServiceMock) UpdatePlantHydration(plantId string, hydration float64) {
	s := ps.hydrationRepo.GetSensorByPlantID(plantId)
	s.Connect() <- sensor.SensorData[float64]{
		PlantID: s.PlantID,
		Data:    hydration,
	}
}

func (ps *PlantsServiceMock) UpdatePlantHealth(plantId string, health plant.HealthData) {
	s := ps.healthRepo.GetSensorByPlantID(plantId)
	s.Connect() <- sensor.SensorData[plant.HealthData]{
		PlantID: s.PlantID,
		Data:    health,
	}
}

func (ps *PlantsServiceMock) UpdatePlantAge(plantId string, age int) {
	s := ps.ageRepo.GetSensorByPlantID(plantId)
	s.Connect() <- sensor.SensorData[int]{
		PlantID: s.PlantID,
		Data:    age,
	}
}