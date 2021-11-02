package handdler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"restfull.com/rest/databasesetup"
)

type Details struct {
	Id       int
	Name     string
	Usn      string
	College  string
	Location string
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
	var mydet []Details
	err := db.Model(&mydet).Select()
	if err != nil {
		fmt.Println(err)
	}
	b, err := json.Marshal(mydet)
	str := string(b)
	fmt.Fprintf(w, str)
	db.Close()
}

func Getadetails(w http.ResponseWriter, r *http.Request) {
	db := databasesetup.Setup()
	vars := mux.Vars(r)
	key := vars["id"]
	intkey, _ := strconv.Atoi(key)
	stud := new(Details)
	err := db.Model(stud).Where("id=?", intkey).Select()
	if err != nil {
		panic(err)
	}
	b, err := json.Marshal(stud)
	str := string(b)
	fmt.Fprintf(w, str)
	db.Close()
}

func Postdetails(w http.ResponseWriter, r *http.Request) {
	db := databasesetup.Setup()
	newpost, _ := ioutil.ReadAll(r.Body)
	var newdet Details
	json.Unmarshal(newpost, &newdet)
	_, err := db.Model(&newdet).Insert()
	if err != nil {
		panic(err)
	}
	Getalldetails(w, r)
	db.Close()
}

func Deletedetails(w http.ResponseWriter, r *http.Request) {
	db := databasesetup.Setup()
	vars := mux.Vars(r)
	val := vars["id"]
	intval, _ := strconv.Atoi(val)
	var det Details
	_, err := db.Model(&det).Where("id=?", intval).Delete()
	if err != nil {
		panic(err)
	}
	Getalldetails(w, r)
	db.Close()
}

func Updatedetails(w http.ResponseWriter, r *http.Request) {
	db := databasesetup.Setup()
	vars := mux.Vars(r)
	val := vars["id"]
	intval, _ := strconv.Atoi(val)
	updateval, _ := ioutil.ReadAll(r.Body)
	var updatedetail Details
	json.Unmarshal(updateval, &updatedetail)
	_, err := db.Model(&updatedetail).Where("id=?", intval).Update()
	if err != nil {
		panic(err)
	}
	Getalldetails(w, r)
	db.Close()
}
