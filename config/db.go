package config

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func GetDb(connURL string) (*sql.DB, error) {
	// connStr := "root:@tcp(127.0.0.1:3306)/ngc15"

	db, err := sql.Open("mysql", connURL)

	return db, err
}
