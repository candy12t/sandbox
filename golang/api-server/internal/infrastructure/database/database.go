package database

import (
	"database/sql"

	"github.com/go-sql-driver/mysql"
)

func NewDB() (*sql.DB, error) {
	conf := mysql.Config{
		User:                 "root",
		Passwd:               "mysql",
		Net:                  "tcp",
		Addr:                 "db:3306",
		DBName:               "api-server",
		AllowNativePasswords: true,
		ParseTime:            true,
	}
	db, err := sql.Open("mysql", conf.FormatDSN())
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
