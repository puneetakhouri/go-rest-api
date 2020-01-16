package main

import "encoding/json"
import "fmt"
import "github.com/gorilla/mux"
import "io/ioutil"
import "log"
import "net/http"

type employee struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

var employees []employee

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/test", test)
	//router.HandleFunc("/addEmployee", addEmployee)
	//router.HandleFunc("/getEmployee/{id}", getEmployee)
	//router.HandleFunc("/getEmployees", getEmployees)
	router.HandleFunc("/employee", addEmployee).Methods("POST")
	router.HandleFunc("/employee", getEmployees).Methods("GET")
	router.HandleFunc("/employee/{id}", getEmployee).Methods("GET")
	log.Fatal(http.ListenAndServe(":9090", router))
}

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func test(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the test url")
}

func addEmployee(w http.ResponseWriter, r *http.Request) {
	var newEmployee employee

	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w, "Kindly enter employee with name and id")
	}

	json.Unmarshal(reqBody, &newEmployee)

	employees = append(employees, newEmployee)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newEmployee)
}

func getEmployee(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	for _, singleEmp := range employees {
		if singleEmp.Id == id {
			json.NewEncoder(w).Encode(singleEmp)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)

}

func getEmployees(w http.ResponseWriter, r *http.Request) {

	json.NewEncoder(w).Encode(employees)

}
