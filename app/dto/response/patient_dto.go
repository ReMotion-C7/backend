package response

type PatientDto struct {
	Id               int    `json:"id"`
	Name             string `json:"name"`
	Phase            string `json:"phase"`
	PhoneNumber      string `json:"phoneNumber"`
	DateOfBirth      string `json:"dateOfBirth"`
	TherapyStartDate string `json:"therapyStartDate"`
}

type PatientDetailDto struct {
	Id                         int                          `json:"id"`
	Name                       string                       `json:"name"`
	Phase                      string                       `json:"phase"`
	Diagnostic                 string                       `json:"diagnostic"`
	PhoneNumber                string                       `json:"phoneNumber"`
	DateOfBirth                string                       `json:"dateOfBirth"`
	Gender                     string                       `json:"gender"`
	TherapyStartDate           string                       `json:"therapyStartDate"`
	Symptoms                   []string                     `json:"symptoms"`
	PatientExerciseForFisioDto []PatientExerciseForFisioDto `json:"exercises"`
	ProgressDto                []ProgressDto                `json:"progresses"`
}

type UserNonFisioDto struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phoneNumber"`
}
