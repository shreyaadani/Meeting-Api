// main project Meeting.go
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Mrec struct {
	Id          string `json:"Id"`
	Title       string `json:"Title"`
	Partcipants string `json:"Partcipants"`
	StartT      string `json:"StartT"`
	EndT        string `json:"EndT"`
	Crtime      string `json:"Crtime"`
}

// let's declare a global Meeting records array
// that we can then populate in our main function
// to simulate a database
var Mrecs []Mrec

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Meeting Management!")
	fmt.Println("Endpoint Hit: Meetings")
}

func handleRequests() {
	// creates a new instance of a mux router
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/all", returnAll)
	myRouter.HandleFunc("/meetings", createNewMeeting).Methods("POST")
	myRouter.HandleFunc("/meetings", GetMeeting).Methods("GET")
	myRouter.HandleFunc("/meeting/{id}", returnSingleMeeting)

	log.Fatal(http.ListenAndServe(":9090", myRouter))
}

func returnAll(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAll Meeting")
	json.NewEncoder(w).Encode(Mrecs)
}

func createNewMeeting(w http.ResponseWriter, r *http.Request) {
	// get the body of our POST request
	// unmarshal this into a new Meeting record struct
	// append this to our Meeting records array.
	reqBody, _ := ioutil.ReadAll(r.Body)
	var mrec Mrec
	json.Unmarshal(reqBody, &mrec)
	// update our global Meeting records array to include
	// our new Meeting record
	mrec.Crtime = time.Now().String()
	Mrecs = append(Mrecs, mrec)

	json.NewEncoder(w).Encode(mrec)
}

func GetMeeting(w http.ResponseWriter, r *http.Request) {

	//keys := r.URL.Query()
	strt := r.URL.Query().Get("StartT")
	endt := r.URL.Query().Get("EndT")
	log.Println("key is", strt, endt)
	log.Println(endt)
	for _, mrec := range Mrecs {
		log.Println(mrec.StartT)
		if mrec.StartT == strt {
			json.NewEncoder(w).Encode(mrec)
			log.Println("if success")
		}
	}
}
func returnSingleMeeting(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	// Loop over all of our Meeting records
	// if the article.Id equals the key we pass in
	// return the article encoded as JSON
	for _, mrec := range Mrecs {
		if mrec.Id == key {
			json.NewEncoder(w).Encode(mrec)
		}
	}
}

func main() {
	fmt.Println("Rest API v2.0 -Shreyaa RestAPI Test")
	Mrecs = []Mrec{
		Mrec{Id: "1", Title: "Hello", Partcipants: "first Description", StartT: "2020-10-20", EndT: "2020-10-10", Crtime: "2020-10-09"},
		Mrec{Id: "2", Title: "Hello 2", Partcipants: "second Description", StartT: "2020-10-10", EndT: "2020-10-10", Crtime: "2020-10-09"},
	}

	handleRequests()
}
