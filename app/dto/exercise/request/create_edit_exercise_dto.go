package request

type CreateExerciseDto struct {
	Name        string `json:"name"`
	TypeId      int    `json:"typeId"`
	Description string `json:"description"`
	Muscle      string `json:"muscle"`
}
