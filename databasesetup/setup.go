package databasesetup

import (
	"fmt"

	"github.com/go-pg/pg"
	_ "github.com/lib/pq"
)

const (
	DB_HOST     = "localhost"
	DB_PORT     = 8080
	DB_USER     = "postgres"
	DB_PASSWORD = "codecraft"
	DB_NAME     = "student"
)

func Setup() *pg.DB {
	/*dinfo := fmt.Sprintf("host=%s port=%d user=%s password=%s name=%s sslmode=disable", DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dinfo)
	if err != nil {
		fmt.Println(err)

	}

	return db*/
	address := fmt.Sprintf("%s:%s", "localhost", "8080")
	dbdetail := &pg.Options{
		User:     "postgres",
		Password: "codecraft",
		Addr:     address,
		Database: "student",
		PoolSize: 50,
	}
	con := pg.Connect(dbdetail)
	if con == nil {
		fmt.Println("Cannot connect")
	}
	return con
}
