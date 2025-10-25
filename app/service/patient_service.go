package service

import (
	"ReMotion-C7/app/dto/request"
	"ReMotion-C7/app/dto/response"
	"ReMotion-C7/app/helper"
	"ReMotion-C7/constant"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func AddPatientService(c *fiber.Ctx, addPatientDto request.AddPatientDto, id int) (error) {

	err := helper.AddPatient(addPatientDto, id)
	if err != nil {
		return err
	}

	return nil

}

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

	if patient.Id == 0 {
		return response.PatientDetailDto{}, fmt.Errorf(constant.ErrPatientNotFound)
	}

	return patient, nil

}

func SearchPatientService(c *fiber.Ctx, patientName string) ([]response.SearchPatientDto, error) {

	patients, err := helper.FindPatientsByName(patientName)
	if err != nil {
		return []response.SearchPatientDto{}, err
	}

	if len(patients) == 0 {
		return []response.SearchPatientDto{}, fmt.Errorf(constant.ErrPatientNotFound)
	}

	return patients, nil

}
