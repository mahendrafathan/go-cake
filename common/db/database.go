package db

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func NewMysqlConnection() (*sql.DB, error) {
	db, err := sql.Open("mysql", os.Getenv("DB_CONN"))
	if err != nil {
		return db, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, err
}
