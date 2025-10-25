package constant

type Success string

var (
	SuccessConnectDatabase Success = "successfully connected to the database."
	SuccessMigrateDatabase Success = "database migration completed successfully."
	SuccessSeedDatabase    Success = "database seeding completed successfully."
	SuccessLoadEnvFile     Success = "successfully loaded .env file."
	SuccessResetTable      Success = "database tables reset successfully."
	SuccessLogin           Success = "successfully login!"
	SuccessRegister        Success = "successfully register!"
)
