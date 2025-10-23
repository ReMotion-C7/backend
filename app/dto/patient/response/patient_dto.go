package response

type PatientDto struct {
	Id               string `json:"id"`
	Name             string `json:"name"`
	Email            string `json:"email"`
	PhoneNumber      string `json:"phoneNumber"`
	DateOfBirth      string `json:"dateOfBirth"`
	Gender           string `json:"gender"`
	TherapyStartDate string `json:"therapyStartDate"`
}
