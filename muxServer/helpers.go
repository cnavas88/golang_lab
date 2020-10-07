package main

import "os"

// DefaultEnv get value of system variable with default value
func DefaultEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}
