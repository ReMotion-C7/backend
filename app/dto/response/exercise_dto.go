package response

type ExerciseDto struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Type        string `json:"type"`
	Description string `json:"description"`
	Muscle      string `json:"muscle"`
	Image       string `json:"image"`
}

type ExerciseDetailForFisioDto struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Type        string `json:"type"`
	Description string `json:"description"`
	Muscle      string `json:"muscle"`
	Image       string `json:"image"`
	Video       string `json:"video"`
}

type PatientExerciseForFisioDto struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Type        string `json:"type"`
	Description string `json:"description"`
	Image       string `json:"image"`
	Muscle      string `json:"muscle"`
	Set         int    `json:"set"`
	RepOrTime   int    `json:"repOrTime"`
}

type ExerciseForPatientDto struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Type        string `json:"type"`
	Description string `json:"description"`
	Image       string `json:"image"`
	Video       string `json:"video"`
	Muscle      string `json:"muscle"`
	Set         int    `json:"set"`
	RepOrTime   int    `json:"repOrTime"`
}
