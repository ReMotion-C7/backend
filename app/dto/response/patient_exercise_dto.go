package response

type PatientExerciseForFisioDto struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Type        string `json:"type"`
	Description string `json:"description"`
	Image       string `json:"image"`
	Muscle      string `json:"muscle"`
	Set         int    `json:"set"`
	RepOrTime   int    `json:"repOrTime"`
}

type ExerciseForPatientDto struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Type      string `json:"type"`
	Video     string `json:"video"`
	Muscle    string `json:"muscle"`
	Set       int    `json:"set"`
	RepOrTime int    `json:"repOrTime"`
}

type ExerciseDetailForPatientDto struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Type        string `json:"type"`
	Description string `json:"description"`
	Video       string `json:"video"`
	Muscle      string `json:"muscle"`
	Set         int    `json:"set"`
	RepOrTime   int    `json:"repOrTime"`
}

type PatientSessionDto struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Type      string `json:"type"`
	Image     string `json:"image"`
	Muscle    string `json:"muscle"`
	Set       int    `json:"set"`
	RepOrTime int    `json:"repOrTime"`
}
