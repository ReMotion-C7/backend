package repository

import (
	"ReMotion-C7/app/dto/request"
	"ReMotion-C7/app/dto/response"
	"ReMotion-C7/app/model"
	"ReMotion-C7/config"
	"ReMotion-C7/constant"
	"fmt"
	"strings"
	"time"
)

func AddPatient(dto request.AddPatientDto, id int) error {

	database := config.GetDatabase()

	therapyStartDate, err := time.Parse("2006-01-02", dto.TherapyStartDate)
	if err != nil {
		return err
	}

	patient := model.Patient{
		UserID:           uint(dto.UserId),
		TherapyStartDate: therapyStartDate,
		Phase:            dto.Phase,
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

	err = database.Where(`patient_id = ?`, patientId).Delete(&model.Symptom{}).Error
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

func RetrievePatients() ([]response.PatientDto, error) {

	database := config.GetDatabase()
	var patients []model.Patient

	err := database.Preload("PatientUser").Find(&patients).Error
	if err != nil {
		return []response.PatientDto{}, err
	}

	var dto []response.PatientDto

	for _, p := range patients {
		dto = append(dto, response.PatientDto{
			Id:               int(p.ID),
			Name:             p.PatientUser.Name,
			PhoneNumber:      p.PatientUser.PhoneNumber,
			Phase:            p.Phase,
			DateOfBirth:      p.PatientUser.DateOfBirth.Format("2006-01-02"),
			TherapyStartDate: p.TherapyStartDate.Format("2006-01-02"),
		})
	}

	return dto, nil

}

func FindPatientsByName(patientName string) ([]response.SearchPatientDto, error) {

	database := config.GetDatabase()
	var patients []model.Patient

	patientName = strings.TrimSpace(patientName)

	err := database.Preload("PatientUser").
		Joins("JOIN users u ON u.id = patients.user_id").
		Where("u.name ILIKE ?", "%"+patientName+"%").
		Find(&patients).Error
	if err != nil {
		return nil, err
	}

	if len(patients) == 0 {
		return nil, fmt.Errorf(constant.ErrPatientNotFound)
	}

	var dto []response.SearchPatientDto
	for _, p := range patients {
		dto = append(dto, response.SearchPatientDto{
			Id:          int(p.ID),
			Name:        p.PatientUser.Name,
			PhoneNumber: p.PatientUser.PhoneNumber,
		})
	}

	return dto, nil

}

func FindPatientPhaseById(patientId int) (int, error) {

	database := config.GetDatabase()

	var patient model.Patient

	err := database.Where(`id = ?`, patientId).First(&patient).Error
	if err != nil {
		return 0, fmt.Errorf(constant.ErrPatientNotFound)
	}

	return patient.Phase, nil

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

	if patient.ID == 0 {
		return response.PatientDetailDto{}, fmt.Errorf(constant.ErrPatientNotFound)
	}

	var symptomNames []string
	for _, s := range patient.Symptoms {
		symptomNames = append(symptomNames, s.Name)
	}

	var dto []response.PatientExerciseForFisioDto
	for _, e := range patient.PatientExercises {
		dto = append(dto, response.PatientExerciseForFisioDto{
			Id:          int(e.ExerciseID),
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
		PatientExerciseForFisioDto: dto,
	}, nil

}
