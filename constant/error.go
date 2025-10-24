package constant

type Error string

const (
	ErrReadEnvFile        = "Error: unable to read .env file."
	ErrConnectingDatabase = "Error: unable to connect to database."
	ErrPingDatabase       = "Error: unable to ping database."
	ErrGetSQLInstance     = "Error: failed to get SQL DB instance."
)
