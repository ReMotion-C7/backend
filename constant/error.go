package constant

const (
	ErrReadEnvFile            = "unable to read .env file."
	ErrConnectingDatabase     = "unable to connect to database."
	ErrPingDatabase           = "unable to ping database."
	ErrGetSQLInstance         = "failed to get SQL DB instance."
	ErrSeedingDatabase        = "unable to seed database."
	ErrFetchDataWhileSeeding  = "unable to fetch data while seeding."
	ErrResetTable             = "unable to reset database tables."
	ErrMigrateDatabase        = "unable to migrate database tables."
	ErrAllInputMustBeFilled   = "all input fields must be filled."
	ErrInvalidCredentials     = "invalid credentials."
	ErrEmailAlreadyRegistered = "email already registered."
	ErrUploadFile             = "error while upload file"
)
