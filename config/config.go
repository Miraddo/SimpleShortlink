package config

import "os"

func getEnv(key string, defaultValue string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return defaultValue
}

var DatabaseHost = getEnv("Shorter_DB_HOST", "127.0.0.1")
var DatabaseUser = getEnv("Shorter_DB_USR", "postgres")
var DatabasePass = getEnv("Shorter_DB_PSS", "123456")
var DatabaseName = getEnv("Shorter_DB_NM", "shortlink")
var DatabaseSSLMode = getEnv("Shorter_DB_SSLMode", "disable")
