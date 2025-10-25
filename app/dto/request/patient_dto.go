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
