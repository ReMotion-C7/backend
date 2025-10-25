package constant

type Success string

var (
	SuccessConnectDatabase Success = "Success: Successfully connected to the database."
	SuccessMigrateDatabase Success = "Success: Database migration completed successfully."
	SuccessSeedDatabase    Success = "Success: Database seeding completed successfully."
	SuccessLoadEnvFile     Success = "Success: Successfully loaded .env file."
	SuccessResetTable      Success = "Success: Database tables reset successfully."
)
