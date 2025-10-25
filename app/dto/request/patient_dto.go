package request

type AddPatientDto struct {
	UserId           int      `json:"userId" validate:"required"`
	Phase            int      `json:"phase" validate:"required"`
	TherapyStartDate string   `json:"therapyStartDate" validate:"required"`
	Symptoms         []string `json:"symptoms" validate:"required"`
}

type EditPatientDto struct {
	Phase    int      `json:"phase" validate:"required"`
	Symptoms []string `json:"symptoms" validate:"required"`
}

type AssignExerciseToPatientDto struct {
	ExerciseId int `json:"exerciseId" validate:"required"`
	Set        int `json:"set" validate:"required"`
	RepOrTime  int `json:"repOrTime" validate:"required"`
}

type EditPatientExerciseDto struct {
	Set       int `json:"set" validate:"required"`
	RepOrTime int `json:"repOrTime" validate:"required"`
}
