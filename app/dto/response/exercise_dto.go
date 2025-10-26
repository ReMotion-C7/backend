package response

type ExerciseDto struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Type        string `json:"type"`
	Description string `json:"description"`
	Muscle      string `json:"muscle"`
	Image       string `json:"image"`
}

type ExerciseModalDto struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Type   string `json:"type"`
	Muscle string `json:"muscle"`
	Image  string `json:"image"`
}

type ExerciseDetailForFisioDto struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Type        string `json:"type"`
	Description string `json:"description"`
	Muscle      string `json:"muscle"`
	Image       string `json:"image"`
	Video       string `json:"video"`
}
