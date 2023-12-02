package utils

import (
	"log"
	"os"
)

func ConnectionString(key string) string {
	connStr, status := os.LookupEnv(key)
	if !status {
		log.Fatalf("Missing environment variable %v", key)
	}

	return connStr
}
