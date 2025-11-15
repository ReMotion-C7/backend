package service

import (
	"ReMotion-C7/app/dto/request"
	"ReMotion-C7/app/dto/response"
	"ReMotion-C7/app/repository"
)

func AddPatientService(dto request.AddPatientDto, id int) error {

	err := repository.AddPatient(dto, id)
	if err != nil {
		return err
	}

	return nil

}

func AddProgressService(dto request.AddProgressDto, id int) error {

	err := repository.AddProgress(dto, id)
	if err != nil {
		return err
	}

	return nil

}

func EditPatientService(dto request.EditPatientDto, fisioId int, patientId int) error {

	err := repository.EditPatient(dto, fisioId, patientId)
	if err != nil {
		return err
	}

	return nil

}

func GetUsersService(fisioId int) ([]response.UserNonFisioDto, error) {

	users, err := repository.RetrieveUsers(fisioId)
	if err != nil {
		return []response.UserNonFisioDto{}, err
	}

	return users, nil

}

func GetPatientsService(fisioId int) ([]response.PatientDto, error) {

	patients, err := repository.RetrievePatients(fisioId)
	if err != nil {
		return nil, err
	}

	return patients, nil

}

func GetPatientProgressesService(patientId int) ([]response.ProgressDto, error) {

	progresses, err := repository.RetrievePatientProgresses(patientId)
	if err != nil {
		return nil, err
	}

	return progresses, nil

}

// func GetPatientPhaseService(patientId int) (int, error) {

// 	phase, err := repository.FindPatientPhaseById(patientId)
// 	if err != nil {
// 		return 0, err
// 	}

// 	return phase, nil

// }

func GetPatientDetail(fisioId int, patientId int) (response.PatientDetailDto, error) {

	patient, err := repository.FindPatientDetail(fisioId, patientId)
	if err != nil {
		return response.PatientDetailDto{}, err
	}

	return patient, nil

}
