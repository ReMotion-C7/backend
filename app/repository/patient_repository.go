package repository

import (
	"ReMotion-C7/app/dto/request"
	"ReMotion-C7/app/dto/response"
	"ReMotion-C7/app/model"
	"ReMotion-C7/config"
	"ReMotion-C7/constant"
	"ReMotion-C7/utils"
	"fmt"
	"time"
)

func AddPatient(dto request.AddPatientDto, id int) error {

	database := config.GetDatabase()

	therapyStartDate, err := time.Parse("2006-01-02", dto.TherapyStartDate)
	if err != nil {
		return err
	}

	if len(dto.Symptoms) == 0 {
		return fmt.Errorf(constant.ErrAllInputMustBeFilled)
	}

	patient := model.Patient{
		UserID:           uint(dto.UserId),
		TherapyStartDate: therapyStartDate,
		Diagnostic:       dto.Diagnostic,
		PhaseID:          utils.PointNumber(dto.PhaseId),
		FisiotherapyID:   uint(id),
	}

	err = database.Create(&patient).Error
	if err != nil {
		return err
	}

	var symptoms []model.Symptom

	for _, s := range dto.Symptoms {

		symptoms = append(symptoms, model.Symptom{
			Name:      s,
			PatientID: patient.ID,
		})

	}

	err = database.Create(&symptoms).Error
	if err != nil {
		return err
	}

	return nil

}

func AddProgress(dto request.AddProgressDto, id int) error {

	database := config.GetDatabase()

	progressDate, err := time.Parse("2006-01-02", dto.Date)
	if err != nil {
		return err
	}

	progress := model.Progress{
		PatientID: uint(id),
		Date:      progressDate,
	}

	err = database.Create(&progress).Error
	if err != nil {
		return err
	}

	return nil
}

func EditPatient(dto request.EditPatientDto, fisioId int, patientId int) error {

	database := config.GetDatabase()

	var patient model.Patient

	err := database.
		Where(`id = ? AND fisiotherapy_id = ?`, patientId, fisioId).
		First(&patient).Error
	if err != nil {
		return fmt.Errorf(constant.ErrPatientNotFound)
	}

	err = database.Model(&patient).Update("phase", dto.Phase).Error
	if err != nil {
		return err
	}

	err = database.Where(`patient_id = ?`, patientId).Unscoped().Delete(&model.Symptom{}).Error
	if err != nil {
		return fmt.Errorf(constant.ErrPatientNotFound)
	}

	var newSymptoms []model.Symptom
	for _, s := range dto.Symptoms {
		newSymptoms = append(newSymptoms, model.Symptom{
			Name:      s,
			PatientID: patient.ID,
		})
	}

	if err := database.Create(&newSymptoms).Error; err != nil {
		return err
	}

	return nil
}

func RetrievePatients(id int) ([]response.PatientDto, error) {

	database := config.GetDatabase()

	var patients []model.Patient

	err := database.Preload("Phase").
		Preload("PatientUser").
		Where(`fisiotherapy_id = ?`, id).
		Find(&patients).Error

	if err != nil {
		return []response.PatientDto{}, err
	}

	var dto []response.PatientDto

	for _, p := range patients {
		dto = append(dto, response.PatientDto{
			Id:               int(p.ID),
			Name:             p.PatientUser.Name,
			PhoneNumber:      p.PatientUser.PhoneNumber,
			Phase:            p.Phase.Name,
			DateOfBirth:      p.PatientUser.DateOfBirth.Format("2006-01-02"),
			TherapyStartDate: p.TherapyStartDate.Format("2006-01-02"),
		})
	}

	return dto, nil

}

func RetrievePatientProgresses(patientId int) ([]response.ProgressDto, error) {

	database := config.GetDatabase()

	var progresses []model.Progress

	err := database.Where(`patient_id = ?`, patientId).Find(&progresses).Error
	if err != nil {
		return []response.ProgressDto{}, err
	}

	var dto []response.ProgressDto

	for _, p := range progresses {
		dto = append(dto, response.ProgressDto{
			Id: int(p.ID),
			Date: p.Date.Format("2006-01-02"),
		})
	}

	return dto, nil

}

func RetrieveUsers(fisioId int) ([]response.UserNonFisioDto, error) {

	database := config.GetDatabase()

	subquery := database.
		Table("patients").
		Select("user_id").
		Where("fisiotherapy_id = ?", fisioId)

	var users []model.User
	err := database.
		Where("role_id = ?", 2).
		Where("id NOT IN (?)", subquery).
		Find(&users).Error
	if err != nil {
		return nil, err
	}

	var dto []response.UserNonFisioDto
	for _, u := range users {
		dto = append(dto, response.UserNonFisioDto{
			Id:          int(u.ID),
			Name:        u.Name,
			PhoneNumber: u.PhoneNumber,
		})
	}

	return dto, nil

}

func FindPatientById(fisioId int, patientId int) error {

	database := config.GetDatabase()

	var patient model.Patient

	err := database.Where(`id = ? AND fisiotherapy_id = ?`, patientId, fisioId).First(&patient).Error
	if err != nil {
		return fmt.Errorf(constant.ErrInvalidPatient)
	}

	return nil

}

// func FindPatientPhaseById(patientId int) (int, error) {

// 	database := config.GetDatabase()

// 	var patient model.Patient

// 	err := database.Where(`id = ?`, patientId).First(&patient).Error
// 	if err != nil {
// 		return 0, fmt.Errorf(constant.ErrPatientNotFound)
// 	}

// 	return patient.Phase, nil

// }

func FindPatientDetail(fisioId int, patientId int) (response.PatientDetailDto, error) {

	database := config.GetDatabase()
	var patient model.Patient

	err := database.
		Preload("PatientExercises.Method").
		Preload("PatientExercises.Exercise").
		Preload("PatientUser.Gender").
		Preload("FisiotherapyUser").
		Preload("Symptoms").
		Preload("Phase").
		Preload("Progresses").
		Where("id = ? AND fisiotherapy_id = ?", patientId, fisioId).Find(&patient).Error

	if err != nil {
		return response.PatientDetailDto{}, err
	}

	if patient.ID == 0 {
		return response.PatientDetailDto{}, fmt.Errorf(constant.ErrPatientNotFound)
	}

	var symptomNames []string
	for _, s := range patient.Symptoms {
		symptomNames = append(symptomNames, s.Name)
	}

	var exercisesDto []response.PatientExerciseForFisioDto
	for _, e := range patient.PatientExercises {

		exerciseId := int(e.ExerciseID)

		exercisesDto = append(exercisesDto, response.PatientExerciseForFisioDto{
			Id:          exerciseId,
			Name:        e.Exercise.Name,
			Method:      e.Method.Name,
			Image:       e.Exercise.Image,
			Muscle:      e.Exercise.Muscle,
			Description: e.Exercise.Description,
			Set:         e.Set,
			RepOrTime:   e.RepOrTime,
		})
	}

	var progressesDto []response.ProgressDto
	for _, p := range patient.Progresses {

		progressesDto = append(progressesDto, response.ProgressDto{
			Id:   int(p.ID),
			Date: p.Date.Format("2006-01-02"),
		})
	}

	return response.PatientDetailDto{
		Id:                         int(patient.ID),
		Name:                       patient.PatientUser.Name,
		Gender:                     patient.PatientUser.Gender.Name,
		PhoneNumber:                patient.PatientUser.PhoneNumber,
		Phase:                      patient.Phase.Name,
		Diagnostic:                 patient.Diagnostic,
		DateOfBirth:                patient.PatientUser.DateOfBirth.Format("2006-01-02"),
		TherapyStartDate:           patient.TherapyStartDate.Format("2006-01-02"),
		Symptoms:                   symptomNames,
		PatientExerciseForFisioDto: exercisesDto,
		ProgressDto:                progressesDto,
	}, nil

}
