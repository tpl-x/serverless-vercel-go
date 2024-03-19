package env

const (
	pgUrlKey = "POSTGRES_URL"
)

// GetPgUrl returns the value of the POSTGRES_URL environment variable
func GetPgUrl() string {
	return getEnv(pgUrlKey)
}
