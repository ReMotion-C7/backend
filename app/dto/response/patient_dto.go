package response

type PatientDto struct {
	Id               string `json:"id"`
	Name             string `json:"name"`
	Phase            int    `json:"phase"`
	PhoneNumber      string `json:"phoneNumber"`
	DateOfBirth      string `json:"dateOfBirth"`
	TherapyStartDate string `json:"therapyStartDate"`
}

type PatientDetailDto struct {
	Id                         string                       `json:"id"`
	Name                       string                       `json:"name"`
	Phase                      int                          `json:"phase"`
	PhoneNumber                string                       `json:"phoneNumber"`
	DateOfBirth                string                       `json:"dateOfBirth"`
	Gender                     string                       `json:"gender"`
	TherapyStartDate           string                       `json:"therapyStartDate"`
	Symptoms                   string                       `json:"symptoms"`
	PatientExerciseForFisioDto []PatientExerciseForFisioDto `json:"exercises"`
}

type SearchPatientDto struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phoneNumber"`
}

type PatientPhaseDto struct {
	Phase int `json:"phase"`
}
