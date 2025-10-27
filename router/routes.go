package router

import (
	"ReMotion-C7/app/controller"
	"ReMotion-C7/app/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetUp(app *fiber.App) {

	api := app.Group("/api/v1")

	// AUTH ENPOINTS
	auth := api.Group("/auth")
	auth.Post("/login", controller.Login)
	auth.Post("/register", controller.Register)

	// EXERCISE ON FISIO ENDPOINTS
	exerciseOnFisio := api.Group("/fisio/exercises", middleware.FisioMiddleware)
	exerciseOnFisio.Post("/create", controller.CreateExercise)
	exerciseOnFisio.Get("/", controller.GetExercises)
	exerciseOnFisio.Get("/modal", controller.GetExercisesModal)
	exerciseOnFisio.Get("/:id", controller.GetExerciseDetail)
	exerciseOnFisio.Delete("/delete/:id", controller.DeleteExercise)

	// PATIENT ON FISIO ENPOINTS
	patientOnFisio := api.Group("/fisio/:id/patients", middleware.FisioMiddleware)
	patientOnFisio.Post("/add", controller.AddPatient)
	patientOnFisio.Get("/", controller.GetPatients)
	patientOnFisio.Patch("/edit/:patient_id", middleware.PatientOwnership, controller.EditPatient)
	patientOnFisio.Get("/:patient_id", middleware.PatientOwnership, controller.GetPatientDetail)

	// PATIENT'S EXERCISE ON FISIO ENPOINTS
	patientsExerciseOnFisio := patientOnFisio.Group("/:patient_id/exercises", middleware.FisioMiddleware, middleware.PatientOwnership)
	patientsExerciseOnFisio.Post("/assign", controller.AssignExercise)
	patientsExerciseOnFisio.Patch("/edit/:exercise_id", controller.EditPatientExercise)
	patientsExerciseOnFisio.Delete("/delete/:exercise_id", controller.DeletePatientExercise)

	// PATIENT'S DETAIL ENPOINTS
	patientsDetail := api.Group("/patients/:id", middleware.UserMiddleware)
	patientsDetail.Get("/sessions", controller.GetPatientSession)
	patientsDetail.Get("/sessions/exercises", controller.GetPatientExercises)
	patientsDetail.Get("/sessions/exercises/:exercise_id", controller.GetPatientExerciseDetail)
	patientsDetail.Get("/phase", controller.GetPatientPhase)

}
