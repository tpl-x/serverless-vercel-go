package env

import "os"

func getEnv(envKey string) string {
	return os.Getenv(envKey)
}
