package config

import (
	"fmt"
	"os"
)

var Configuration = DB()

func DB() Db {
	return Db{
		PublicHost: gETEnv("PUBLIC_HOST", "https://localhost"),
		Port:       gETEnv("PORT", "8080"),
		DBUser:     gETEnv("DBUSER", "**********"),
		DBAddr:     fmt.Sprintf("%s:%s", gETEnv("DB_HOST", "***************"), gETEnv("DB_POST", "***************")),
		DBPassword: gETEnv("DBPASSWORD", "**********"),
		DBName:     gETEnv("DBNAME", "data12"),
	}
}
func gETEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
