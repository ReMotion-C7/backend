package helper

import (
	"ReMotion-C7/app/dto/response"
	"ReMotion-C7/app/model"
	"ReMotion-C7/config"
)

func RetrievePatients() ([]response.PatientDto, error) {

	database := config.GetDatabase()
	var patients []model.Patient

	err := database.Preload("User").Find(&patients).Error
	if err != nil {
		return []response.PatientDto{}, err
	}

	var patientsDto []response.PatientDto

	for _, p := range patients {
		patientsDto = append(patientsDto, response.PatientDto{
			Id:               int(p.ID),
			Name:             p.PatientUser.Name,
			PhoneNumber:      p.PatientUser.PhoneNumber,
			Phase:            p.Phase,
			DateOfBirth:      p.PatientUser.DateOfBirth.Format("2006-01-02"),
			TherapyStartDate: p.TherapyStartDate.Format("2006-01-02"),
		})
	}

	return patientsDto, nil

}

func FindPatientDetail(fisioId int, patientId int) (response.PatientDetailDto, error) {

	database := config.GetDatabase()
	var patient model.Patient

	err := database.
		Preload("PatientUser.Gender").
		Preload("FisiotherapyUser").
		Preload("FisiotherapyUser").
		Preload("Symptoms").
		Preload("PatientExercises.Exercise.Type").
		Where("id = ? AND fisiotherapy_id = ?", patientId, fisioId).Find(&patient).Error
	if err != nil {
		return response.PatientDetailDto{}, err
	}

	var symptomNames []string
	for _, s := range patient.Symptoms {
		symptomNames = append(symptomNames, s.Name)
	}

	var patientExercises []response.PatientExerciseForFisioDto
	for _, e := range patient.PatientExercises {
		patientExercises = append(patientExercises, response.PatientExerciseForFisioDto{
			Id:          int(e.ID),
			Name:        e.Exercise.Name,
			Type:        e.Exercise.Type.Name,
			Image:       e.Exercise.Image,
			Muscle:      e.Exercise.Muscle,
			Description: e.Exercise.Description,
			Set:         e.Set,
			RepOrTime:   e.RepOrTime,
		})
	}

	return response.PatientDetailDto{
		Id:                         int(patient.ID),
		Name:                       patient.PatientUser.Name,
		Gender:                     patient.PatientUser.Gender.Name,
		PhoneNumber:                patient.PatientUser.PhoneNumber,
		Phase:                      patient.Phase,
		DateOfBirth:                patient.PatientUser.DateOfBirth.Format("2006-01-02"),
		TherapyStartDate:           patient.TherapyStartDate.Format("2006-01-02"),
		Symptoms:                   symptomNames,
		PatientExerciseForFisioDto: patientExercises,
	}, nil

}
