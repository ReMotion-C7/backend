package request

type AddPatientDto struct {
	UserId           int      `json:"userId" validate:"required"`
	PhaseId          int      `json:"phaseId" validate:"required"`
	TherapyStartDate string   `json:"therapyStartDate" validate:"required"`
	Diagnostic       string   `json:"diagnostic" validate:"required"`
	Symptoms         []string `json:"symptoms" validate:"required"`
}

type EditPatientDto struct {
	Phase    int      `json:"phase" validate:"required"`
	Symptoms []string `json:"symptoms" validate:"required"`
}
