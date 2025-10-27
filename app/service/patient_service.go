package service

import (
	"ReMotion-C7/app/dto/request"
	"ReMotion-C7/app/dto/response"
	"ReMotion-C7/app/repository"

	"github.com/gofiber/fiber/v2"
)

func AddPatientService(dto request.AddPatientDto, id int) error {

	err := repository.AddPatient(dto, id)
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

func GetPatientsService(c *fiber.Ctx, fisioId int) (interface{}, error) {

	patientName := c.Query("name")

	if patientName == "" {

		patients, err := repository.RetrievePatients(fisioId)
		if err != nil {
			return nil, err
		}

		return patients, nil

	}

	patients, err := repository.FindPatientsByName(patientName)
	if err != nil {
		return nil, err
	}

	return patients, nil

}

func GetPatientPhaseService(patientId int) (int, error) {

	phase, err := repository.FindPatientPhaseById(patientId)
	if err != nil {
		return 0, err
	}

	return phase, nil

}

func GetPatientDetail(fisioId int, patientId int) (response.PatientDetailDto, error) {

	patient, err := repository.FindPatientDetail(fisioId, patientId)
	if err != nil {
		return response.PatientDetailDto{}, err
	}

	return patient, nil

}
