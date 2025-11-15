package request

type AssignExerciseToPatientDto struct {
	ExerciseId int `json:"exerciseId" validate:"required"`
	MethodId   int `json:"methodId" validate:"required"`
	Set        int `json:"set" validate:"required"`
	RepOrTime  int `json:"repOrTime" validate:"required"`
}

type EditPatientExerciseDto struct {
	Set       int `json:"set" validate:"required"`
	RepOrTime int `json:"repOrTime" validate:"required"`
}
