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
		DBUser:     gETEnv("DBUSER", "root"),
		DBAddr:     fmt.Sprintf("%s:%s", gETEnv("DB_HOST", "127.0.0.1"), gETEnv("DB_POST", "3306")),
		DBPassword: gETEnv("DBPASSWORD", "0987poiulkjh"),
		DBName:     gETEnv("DBNAME", "data12"),
	}
}
func gETEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
