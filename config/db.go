package config

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
)

func MYSQL(my mysql.Config) (*sql.DB, error) {
	db, err := sql.Open("mysql", my.FormatDSN())
	if err != nil {
		return nil, err
	}
	return db, nil
}
