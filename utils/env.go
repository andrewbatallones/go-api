package utils

import "os"

// Gets the ENV variable or returns the provided default value.
func GetEnv(key, defaultVal string) string {
	env, ok := os.LookupEnv(key)
	if !ok {
		env = defaultVal
	}

	return env
}
