package env

import "os"

func SqlSource() string {
	return os.Getenv("SQL_SOURCE")
}

func IsDevelopment() bool {
	return os.Getenv("APP_ENVIRONMENT") == "development"
}

func IsStaging() bool {
	return os.Getenv("APP_ENVIRONMENT") == "staging"
}

func IsProduction() bool {
	return os.Getenv("APP_ENVIRONMENT") == "production"
}

func GetPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return port
}
