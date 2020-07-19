package env

import "os"

func SqlSource() string {
	return os.Getenv("SQL_SOURCE")
}
