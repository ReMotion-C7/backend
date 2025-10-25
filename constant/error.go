package constant

type Error string

const (
	ErrReadEnvFile           = "Error: unable to read .env file."
	ErrConnectingDatabase    = "Error: unable to connect to database."
	ErrPingDatabase          = "Error: unable to ping database."
	ErrGetSQLInstance        = "Error: failed to get SQL DB instance."
	ErrSeedingDatabase       = "Error: unable to seed database."
	ErrFetchDataWhileSeeding = "Error: unable to fetch data while seeding."
	ErrResetTable            = "Error: unable to reset database tables."
	ErrMigrateDatabase       = "Error: unable to migrate database tables."
)
