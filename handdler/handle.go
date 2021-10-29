package handdler

import (
	"fmt"
	"net/http"

	"restfull.com/rest/databasesetup"
)

type Details struct {
	tableName struct{} `pg:"public.students"` // default values are the same
	Id        int      `pg:",pk"`
	Name      string
	Usn       string
	College   string
	Location  string
}

func Welcomeresponse(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome")
}

/*func SelectDBPost(pg *pg.DB, post Details) (Details, error) {
	err := pg.Model(&post).Where("id = 1").Select()
	return post, err
}*/
func Getalldetails(w http.ResponseWriter, r *http.Request) {
	db := databasesetup.Setup()
	//defer db.Close()
	fmt.Println("Print student")
	//val, err := SelectDBPost(db, det)
	//_, err := db.Exec("SELECT * FROM students")
	/*rows, err := db.Query("SELECT * FROM students")*/
	/*if err != nil {
		panic(err)
	}*/
	det := new(Details)
	err := db.Model(&det).Where("id=?", 1).Select()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(det)
}
