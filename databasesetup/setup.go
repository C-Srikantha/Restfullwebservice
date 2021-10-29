package databasesetup

import (
	"database/sql"
	"fmt"
)

const (
	DB_USER     = "Postgresql"
	DB_PASSWORD = "codecraft"
	DB_NAME     = "student"
)

func Setup() *sql.DB {
	dinfo := fmt.Sprintf("user=%s password=%s name=%s", DB_USER, DB_PASSWORD, DB_NAME)
	db, _ := sql.Open("Postgresql", dinfo)
	return db
}
