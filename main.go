package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/adraynrion/ac-server/rest"
	"github.com/gorilla/mux"
)

var Masters []rest.Master

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: homePage")

	fmt.Fprintf(w, "Welcome to the homepage!")
}

func returnAllMasters(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllMasters")

	json.NewEncoder(w).Encode(Masters)
}

func returnSingleMaster(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnSingleMaster")

	vars := mux.Vars(r)
	id := vars["id"]

	for _, m := range Masters {
		if uId, err := strconv.ParseUint(id, 10, 64); err == nil {
			if m.Id == uId {
				json.NewEncoder(w).Encode(m)
				return
			}
		}
	}

	fmt.Fprintf(w, "No master find with given id!")
}

func addNewMaster(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: addNewMaster")

	if reqBody, err := ioutil.ReadAll(r.Body); err == nil {
		var master rest.Master

		json.Unmarshal(reqBody, &master)
		Masters = append(Masters, master)

		json.NewEncoder(w).Encode(master)
	}
}

func freeMaster(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: freeMaster")

	vars := mux.Vars(r)
	id := vars["id"]

	for index, m := range Masters {
		if uId, err := strconv.ParseUint(id, 10, 64); err == nil {
			if m.Id == uId {
				Masters = append(Masters[:index], Masters[index+1:]...)
				json.NewEncoder(w).Encode(m)
				return
			}
		}
	}
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/masters", returnAllMasters)
	myRouter.HandleFunc("/master", addNewMaster).Methods(http.MethodPost)
	myRouter.HandleFunc("/master/{id}", freeMaster).Methods(http.MethodDelete)
	myRouter.HandleFunc("/master/{id}", returnSingleMaster)

	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func main() {
	fmt.Println("Rest API v2.0 - Mux Routers")
	Masters = []rest.Master{
		{Id: 12, Name: "Master 12", Vassals: []rest.Vassal{
			{Id: 98, Name: "Vassal 98", WorkSpeed: 8},
		}},
		{Id: 23, Name: "Master 23", Vassals: []rest.Vassal{
			{Id: 87, Name: "Vassal 87", WorkSpeed: 5},
			{Id: 76, Name: "Vassal 76", WorkSpeed: 3},
		}},
		{Id: 34, Name: "Master 34", Vassals: []rest.Vassal{}},
	}
	handleRequests()
}
