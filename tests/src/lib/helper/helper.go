package helper

import "os"

func GetEnv(key string, fallback string) string {
	envVariable := os.Getenv(key)
	if envVariable == "" {
		return fallback
	}
	return envVariable
}
