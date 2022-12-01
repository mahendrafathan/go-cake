package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func NewMysqlConnection() (*sql.DB, error) {
	db, err := sql.Open("mysql", "user:user@tcp(full_db_mysql:3306)/mydb?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		return db, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, err
}
