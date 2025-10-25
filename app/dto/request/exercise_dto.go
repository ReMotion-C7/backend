package request

type CreateEditExerciseDto struct {
	Name        string `json:"name" validate:"required"`
	TypeId      int    `json:"typeId" validate:"required"`
	Description string `json:"description" validate:"required"`
	Muscle      string `json:"muscle" validate:"required"`
}