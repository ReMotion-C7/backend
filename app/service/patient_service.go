package service

import (
	"ReMotion-C7/app/dto/response"
	"ReMotion-C7/app/helper"

	"github.com/gofiber/fiber/v2"
)

func GetPatientsService(c *fiber.Ctx) ([]response.PatientDto, error) {

	patients, err := helper.RetrievePatients()
	if err != nil {
		return []response.PatientDto{}, err
	}

	return patients, nil

}

func GetPatientDetail(c *fiber.Ctx, fisioId int, patientId int) (response.PatientDetailDto, error) {

	patient, err := helper.FindPatientDetail(fisioId, patientId)
	if err != nil {
		return response.PatientDetailDto{}, err
	}

	return patient, nil

}
