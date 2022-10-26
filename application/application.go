package application

import (
	"context"
	"fmt"
	droneRepository "github.com/SchoolGolang/multithreading-practice/drone/repository"
	"github.com/SchoolGolang/multithreading-practice/listener"
	"github.com/SchoolGolang/multithreading-practice/mock"
	"github.com/SchoolGolang/multithreading-practice/plant"
	"github.com/SchoolGolang/multithreading-practice/plant/repository"
	"github.com/SchoolGolang/multithreading-practice/processor"
	"github.com/SchoolGolang/multithreading-practice/sensor"
	sensorRepo "github.com/SchoolGolang/multithreading-practice/sensor/repository"
	"sync"
	"time"
)

func Run(ctx context.Context) {
	plantsRepo := repository.NewRepository()

	phSensorsRepo := sensorRepo.NewRepository[int]()
	hydrationSensorsRepo := sensorRepo.NewRepository[float64]()
	healthSensorsRepo := sensorRepo.NewRepository[plant.HealthData]()
	ageSensorsRepo := sensorRepo.NewRepository[int]()
	plantsService := mock.NewPlantsServiceMock(plantsRepo, phSensorsRepo, hydrationSensorsRepo, healthSensorsRepo,ageSensorsRepo, 20)

	for i := 0; i < 10; i++ {
		plantsService.AddPlant()
	}

	phListener := listener.NewListener[sensor.SensorData[int]]()
	hydrationListener := listener.NewListener[sensor.SensorData[float64]]()
	healthListener := listener.NewListener[sensor.SensorData[plant.HealthData]]()
	ageListener := listener.NewListener[sensor.SensorData[int]]()

	for _, s := range phSensorsRepo.GetAll() {
		phListener.AddChan(s.Connect())
	}

	for _, s := range hydrationSensorsRepo.GetAll() {
		hydrationListener.AddChan(s.Connect())
	}

	for _, s := range healthSensorsRepo.GetAll() {
		healthListener.AddChan(s.Connect())
	}
	for _, s := range ageSensorsRepo.GetAll() {
		ageListener.AddChan(s.Connect())
	}
	go plantsService.SendRandomUpdates(ctx)

	phOut := phListener.Listen(ctx)
	hydrationOut := hydrationListener.Listen(ctx)
	healthOut := healthListener.Listen(ctx)
	ageOut := ageListener.Listen(ctx)

	dronesRepo := droneRepository.NewDroneRepo(plantsService, plantsRepo)

	phProcessor := processor.NewPHProcessor(plantsRepo, phOut, dronesRepo)
	go phProcessor.RunProcessor(ctx)

	hydrationProcessor := processor.NewHydrationProcessor(plantsRepo, hydrationOut, dronesRepo)
	go hydrationProcessor.RunProcessor(ctx)

	healthProcessor := processor.NewHealthProcessor(plantsRepo, healthOut, dronesRepo)
	go healthProcessor.RunProcessor(ctx)

	ageProcessor := processor.NewAgeProcessor(plantsRepo, ageOut , dronesRepo)
	go ageProcessor.RunProcessor(ctx)

	wg := sync.WaitGroup{}
	wg.Add(1)
	go gracefulShutdown(ctx, &wg, phSensorsRepo, hydrationSensorsRepo, healthSensorsRepo,ageSensorsRepo)
	wg.Wait()

}

func gracefulShutdown(
	ctx context.Context,
	wg *sync.WaitGroup,
	phSensorsRepo *sensorRepo.SensorRepo[int],
	hydrationSensorsRepo *sensorRepo.SensorRepo[float64],
	healthSensorsRepo *sensorRepo.SensorRepo[plant.HealthData],
	ageSensorsRepo *sensorRepo.SensorRepo[int],

) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("graceful shutdown, to force press Ctrl+C")
			for _, s := range phSensorsRepo.GetAll() {
				s.Disconnect()
			}
			for _, s := range hydrationSensorsRepo.GetAll() {
				s.Disconnect()
			}
			for _, s := range healthSensorsRepo.GetAll() {
				s.Disconnect()
			}
			for _, s := range ageSensorsRepo.GetAll() {
				s.Disconnect()
			}

			fmt.Println("Gracefully stopped")
			wg.Done()
			return
		default:
			time.Sleep(1000 * time.Millisecond)
		}
	}
}
