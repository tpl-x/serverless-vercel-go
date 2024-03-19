package env

const (
	kvUrlKey   = "KV_REST_API_URL"
	kvTokenKey = "KV_REST_API_TOKEN"
)

// GetKvUrl returns the value of the KV_REST_API_URL environment variable
func GetKvUrl() string {
	return getEnv(kvUrlKey)
}

// GetKvToken returns the value of the KV_REST_API_TOKEN environment variable
func GetKvToken() string {
	return getEnv(kvTokenKey)
}
