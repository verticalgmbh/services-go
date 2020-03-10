package cfg

import "os"

// EnvConfig get configuration value from environment
func EnvConfig(key string, defaultvalue string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		value = defaultvalue
	}

	return value
}
