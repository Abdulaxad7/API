package config

import (
	"ALBUMS/config/config"
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
)

func Connect() (string, error) {
	_, err := getConnection()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Configuration.DBUser, config.Configuration.DBPassword, config.Configuration.DBAddr, config.Configuration.DBName), err
}

func getConnection() (*sql.DB, error) {
	return MYSQL(mysql.Config{
		User:                 config.Configuration.DBUser,
		Passwd:               config.Configuration.DBPassword,
		Addr:                 config.Configuration.DBAddr,
		DBName:               config.Configuration.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})
}
