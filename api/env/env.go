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
	return os.Getenv("PORT")
}
