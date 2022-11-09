package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Tickets struct {
	Tickets [][2]string `json:"tickets"`
}

type Route struct {
	Route [2]string `json:"route"`
}

func calculation(tickets Tickets) (string, string) {
	var start string
	var end string

	for _, ticket1 := range tickets.Tickets {
		startFlag := false
		endFlag := false
		for _, ticket2 := range tickets.Tickets {
			if ticket1[0] == ticket2[1] {
				startFlag = true
			}
			if ticket1[1] == ticket2[0] {
				endFlag = true
			}
		}
		if startFlag == false {
			start = ticket1[0]
		}
		if endFlag == false {
			end = ticket1[1]
		}
	}
	return start, end
}

func findRoute(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var tickets Tickets
	json.Unmarshal(reqBody, &tickets)

	// Main calculation
	start, end := calculation(tickets)

	var route Route
	route.Route[0] = start
	route.Route[1] = end

	json.NewEncoder(w).Encode(route)
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/route", findRoute).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", myRouter))
}
func main() {
	fmt.Println("Simple API for flight")

	handleRequests()
}
