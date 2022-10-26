package repository

import "github.com/SchoolGolang/multithreading-practice/plant"

type PlantsRepository struct {
	plantsData map[string]plant.PlantData
}

func NewRepository() *PlantsRepository {
	return &PlantsRepository{
		plantsData: make(map[string]plant.PlantData),
	}
}

func (pr *PlantsRepository) AddPlant(data plant.PlantData) {
	if _, ok := pr.plantsData[data.ID]; ok {
		return
	}

	pr.plantsData[data.ID] = data
}

func (pr *PlantsRepository) RemovePlant(id string) {
	delete(pr.plantsData, id)
}

func (pr *PlantsRepository) GetPlant(id string) plant.PlantData {
	return pr.plantsData[id]
}

func (pr *PlantsRepository) GetHydration(id string) float64 {
	return pr.plantsData[id].CurrentHydration
}

func (pr *PlantsRepository) GetPh(id string) int {
	return pr.plantsData[id].CurrentPh
}

func (pr *PlantsRepository) GetHealth(id string) plant.HealthData {
	return pr.plantsData[id].CurrentHealth
}

func (pr *PlantsRepository) GetNormalHydration(id string) float64 {
	return pr.plantsData[id].NormalHydration
}

func (pr *PlantsRepository) GetNormalPh(id string) (int, int) {
	plnt := pr.plantsData[id]
	return plnt.NormalLowerPh, plnt.NormalUpperPh
}

func (pr *PlantsRepository) GetPlantIds() []string {
	ids := make([]string, 0, len(pr.plantsData))
	for k := range pr.plantsData {
		ids = append(ids, k)
	}

	return ids
}

func (pr *PlantsRepository) GetAge(id string) int {
	return pr.plantsData[id].Age
}

func (pr *PlantsRepository) SetPh(id string, ph int) {
	if p, ok := pr.plantsData[id]; ok {
		p.CurrentPh = ph
	}

}

func (pr *PlantsRepository) SetHydration(id string, hydration float64) {
	if p, ok := pr.plantsData[id]; ok {
		p.CurrentHydration = hydration
	}
}

func (pr *PlantsRepository) SetHealth(id string, data plant.HealthData) {
	if p, ok := pr.plantsData[id]; ok {
		p.CurrentHealth = data
	}
}
