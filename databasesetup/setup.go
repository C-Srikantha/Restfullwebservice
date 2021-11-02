package databasesetup

import (
	"fmt"

	"github.com/go-pg/pg"
	_ "github.com/lib/pq"
)

func Setup() *pg.DB {
	//address := fmt.Sprintf("%s:%s", "localhost", "8080")
	dbdetail := &pg.Options{
		User:     "postgres",
		Password: "codecraft",
		Addr:     ":8080",
		Database: "mydatabase",
	}
	con := pg.Connect(dbdetail)
	if con == nil {
		fmt.Println("Cannot connect")
	}
	/*ctx := context.Background()

	_, err := con.ExecContext(ctx, "SELECT 1")
	if err != nil {
		panic(err)
	}*/
	return con
}
