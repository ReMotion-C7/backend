package constant

type Error string

const (
	ErrReadEnvFile           = "unable to read .env file."
	ErrConnectingDatabase    = "unable to connect to database."
	ErrPingDatabase          = "unable to ping database."
	ErrGetSQLInstance        = "failed to get SQL DB instance."
	ErrSeedingDatabase       = "unable to seed database."
	ErrFetchDataWhileSeeding = "unable to fetch data while seeding."
	ErrResetTable            = "unable to reset database tables."
	ErrMigrateDatabase       = "unable to migrate database tables."
	ErrAllInputMustBeFilled  = "all input fields must be filled."
	ErrInvalidCredentials    = "invalid credentials."
)
