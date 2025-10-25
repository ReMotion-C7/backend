package router

import (
	"ReMotion-C7/app/controller"

	"github.com/gofiber/fiber/v2"
)

func SetUp(app *fiber.App) {

	api := app.Group("/api/v1")

	// AUTH ENPOINTS
	auth := api.Group("/auth")
	auth.Post("/login", controller.Login)
	auth.Post("/register", controller.Register)

	// EXERCISE ON FISIO ENDPOINTS
	exerciseOnFisio := api.Group("/fisio/exercises")
	exerciseOnFisio.Post("/create", controller.CreateExercise)
	exerciseOnFisio.Get("/", controller.GetExercises)
	exerciseOnFisio.Get("/:id", controller.GetExerciseDetail)

	// PATIENT ON FISIO ENPOINTS
	patientOnFisio := api.Group("/fisio/:id/patients")
	patientOnFisio.Post("/add", controller.AddPatient)
	patientOnFisio.Patch("/edit/:patient_id", controller.EditPatient)
	patientOnFisio.Get("/", controller.GetPatients)
	patientOnFisio.Get("/:patient_id", controller.GetPatientDetail)

	// PATIENT'S EXERCISE ON FISIO ENPOINTS
	patientsExerciseOnFisio := api.Group("/fisio/:id/patients/:patient_id/exercises")
	patientsExerciseOnFisio.Post("/assign", controller.AssignExercise)
	patientsExerciseOnFisio.Patch("/edit/:exercise_id", controller.EditPatientExercise)

}
