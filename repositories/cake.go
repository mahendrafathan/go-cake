package repositories

import "database/sql"

func NewCakeRepo(db *sql.DB) cakeRepoMysql {
	return cakeRepoMysql{
		db: db,
	}
}
