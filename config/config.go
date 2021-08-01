package config

import "os"

func getEnv(key string, defaultValue string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return defaultValue
}

var DBHost = getEnv("Shorter_DB_HOST", "127.0.0.1")
var DBUser = getEnv("Shorter_DB_USR", "postgres")
var DBPass = getEnv("Shorter_DB_PSS", "123456")
var DBName = getEnv("Shorter_DB_NM", "shortlink")
var DBSSLMode = getEnv("Shorter_DB_SSLMode", "disable")
