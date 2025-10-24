package request

type AddPatientDto struct {
	PatientId        int      `json:"patientId"`
	Phase            int      `json:"phase"`
	TherapyStartDate string   `json:"therapyStartDate"`
	Symptoms         []string `json:"symptoms"`
}

type EditPatientDto struct {
	Phase    int      `json:"phase"`
	Symptoms []string `json:"symptoms"`
}

type AssignExerciseToPatientDto struct {
	ExerciseId int `json:"exerciseId"`
	Set        int `json:"set"`
	RepOrTime  int `json:"repOrTime"`
}

type EditPatientExerciseDto struct {
	Set       int `json:"set"`
	RepOrTime int `json:"repOrTime"`
}
