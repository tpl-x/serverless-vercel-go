package env

import "os"

func GetEnv(envKey string) string {
	return os.Getenv(envKey)
}
