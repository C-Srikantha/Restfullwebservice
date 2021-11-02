package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"restfull.com/rest/handdler"
)

type Details struct {
	Id       int
	Name     string
	Usn      string
	College  string
	Location string
}

var m []Details

func homepage(w http.ResponseWriter, r *http.Request) {
	//fmt.Println(r.Method)
	fmt.Fprintln(w, "Welcome")
}
func getalldetails(w http.ResponseWriter, r *http.Request) { //this function will retrieve all the student data
	fmt.Println(r.Method)
	b, err := json.Marshal(m) //to encode into json
	str := string(b)          //converts byte to string
	if err != nil {
		fmt.Fprintf(w, "ERROR")
	} else {
		fmt.Fprintf(w, str)
	}

}
func getadetail(w http.ResponseWriter, r *http.Request) { //this function will retrieve data of requested student
	vars := mux.Vars(r) //returs a map
	key := vars["id"]
	inkey, _ := strconv.Atoi(key) //converts string to int
	for _, val := range m {
		if val.Id == inkey {
			b, err := json.Marshal(val) //decodeds to json and return byte and a error
			str := string(b)            //converts byte to string
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Fprintf(w, str)
			}
		}

	}

}

func postdetails(w http.ResponseWriter, r *http.Request) { //it will add the new student details

	detail, _ := ioutil.ReadAll(r.Body) //reads r and returns data in byte
	//fmt.Fprintf(w, "%v", string(detail))
	var addnewuser Details              //creating a local var of type Details
	json.Unmarshal(detail, &addnewuser) //decodes json to struct format
	for _, val := range m {
		if val.Id == addnewuser.Id {
			fmt.Fprintf(w, "Id already Present")
			return
		}

	}
	m = append(m, addnewuser)

	b, err := json.Marshal(m) //encode struct to json
	str := string(b)          //convert byte to string
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Fprintf(w, str)
	}

}
func deletedetails(w http.ResponseWriter, r *http.Request) { //it deletes the requeted student details
	vars := mux.Vars(r)           //func Vars returns route varible for current request
	key := vars["id"]             //		returs value of the key id
	inkey, _ := strconv.Atoi(key) //converts string to int
	for index, val := range m {
		if val.Id == inkey {
			m = append(m[:index], m[index+1:]...) //deletes the index of current request id
		}
	}
	b, err := json.Marshal(m) //encode struct to json
	str := string(b)          //convert byte to string
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Fprintf(w, str)
	}

}
func updatedetails(w http.ResponseWriter, r *http.Request) { //it updates the entire details of requested student
	vars := mux.Vars(r)                 //func Vars returns route varible for current request
	key := vars["id"]                   //returns a value of the key mentioned
	val, _ := strconv.Atoi(key)         //converts string to int
	detail, _ := ioutil.ReadAll(r.Body) //reads r and returns the byte
	var updateddetail Details
	json.Unmarshal(detail, &updateddetail) //decodes
	for index, d := range m {
		if d.Id == val {
			d.Id = updateddetail.Id
			d.Name = updateddetail.Name
			d.Usn = updateddetail.Usn
			d.College = updateddetail.College
			d.Location = updateddetail.Location
			m[index] = d
		}

	}
	b, err := json.Marshal(m) //encode struct to json
	str := string(b)          //convert byte to string
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Fprintf(w, str)
	}

}

/*func handlerequest() {
	http.HandleFunc("/", homepage)
	http.HandleFunc("/details", writejson)
	error := http.ListenAndServe(":8080", nil)
	fmt.Println(error)
}*/
func updateadetails(w http.ResponseWriter, r *http.Request) { //it update some feilds of requested student
	vars := mux.Vars(r)                 //it returns route variable for current request
	val := vars["id"]                   //access value of id
	inval, _ := strconv.Atoi(val)       //converts string to int
	detail, _ := ioutil.ReadAll(r.Body) //reads the body till EOF and returs byte
	var det Details
	json.Unmarshal(detail, &det) //decodes json to struct
	for index, d := range m {
		if inval == d.Id {
			if det.Id != 0 {
				d.Id = det.Id
			}
			if det.Name != "" {
				d.Name = det.Name
			}
			if det.Usn != "" {
				d.Usn = det.Usn
			}
			if det.College != "" {
				d.College = det.College
			}
			if det.Location != "" {
				d.Location = det.Location
			}
			m[index] = d

		}
	}
	b, err := json.Marshal(m) //encode struct to json
	str := string(b)          //convert byte to string
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Fprintf(w, str)
	}

}
func handlerequest() {
	myrouter := mux.NewRouter() //initializing a router

	myrouter.HandleFunc("/", homepage) //it registers the Api signature and the function to the mux

	myrouter.HandleFunc("/alldetails", getalldetails).Methods("GET") //it registers the Api signature and the function to the mux

	myrouter.HandleFunc("/adddetails", postdetails).Methods("POST") //it registers the Api signature and the function to the mux

	myrouter.HandleFunc("/details/{id}", getadetail).Methods("GET") //it registers the Api signature and the function to the mux

	myrouter.HandleFunc("/deletedetails/{id}", deletedetails).Methods("DELETE") //it registers the Api signature and the function to the mux

	myrouter.HandleFunc("/updatedetails/{id}", updatedetails).Methods("PUT") //it registers the Api signature and the function to the mux

	myrouter.HandleFunc("/updatedetails/{id}", updateadetails).Methods("PATCH") //it registers the Api signature and the function to the mux

	log.Fatal(http.ListenAndServe(":8081", myrouter)) //creates a server and assign a port and router,it listens to port and send the request to mux
}
func handlerequest_database() {
	router := mux.NewRouter()
	router.HandleFunc("/", handdler.Welcomeresponse).Methods("GET")                    //it registers the Api signature and the function to the mux
	router.HandleFunc("/getdetails", handdler.Getalldetails).Methods("GET")            //it registers the Api signature and the function to the mux
	router.HandleFunc("/getdetails/{id}", handdler.Getadetails).Methods("GET")         //it registers the Api signature and the function to the mux
	router.HandleFunc("/adddetails", handdler.Postdetails).Methods("POST")             //it registers the Api signature and the function to the mux
	router.HandleFunc("/deletedetails/{id}", handdler.Deletedetails).Methods("DELETE") //it registers the Api signature and the function to the mux
	router.HandleFunc("/updatedetails/{id}", handdler.Updatedetails).Methods("PUT")    //it registers the Api signature and the function to the mux
	log.Fatal(http.ListenAndServe(":8081", router))                                    //creates a server and assign a port and router,it listens to port and send the request to mux
}
func main() {
	m = []Details{
		{1, "Srikantha", "4Cb17Cs022", "CEC", "Mangalore"},
		{2, "Adarsh", "4cb17cs023", "AJ", "Udupi"},
		{3, "Anil", "4cb17cs007", "Nitte", "Bangalore"},
		{4, "Roy", "4CB17CS013", "RV", "Bangalore"},
	}
	//handlerequest()
	handlerequest_database()
}
