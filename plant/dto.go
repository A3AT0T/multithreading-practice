package plant

type PlantData struct {
	ID               string
	Name             string
	NormalHydration  float64
	NormalLowerPh    int
	NormalUpperPh    int
	CurrentHydration float64
	CurrentPh        int
	CurrentHealth    HealthData
	Age				 int
}

type HealthData struct {
	LeavesState float64
	RootsState  float64
}
