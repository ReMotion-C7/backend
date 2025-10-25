package request

type CreateEditExerciseDto struct {
	Name        string `json:"name" validate:"required"`
	TypeId      int    `json:"typeId" validate:"required"`
	Description string `json:"description" validate:"required"`
	Muscle      string `json:"muscle" validate:"required"`
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