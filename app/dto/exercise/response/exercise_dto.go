package response

type ExerciseDto struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Type        string `json:"type"`
	Description string `json:"description"`
	Muscle      string `json:"muscle"`
	Image       string `json:"image"`
	Video       string `json:"video"`
}
