package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	// waiting for requests on port 3000
	router := mux.NewRouter()
	router.HandleFunc("/signup", PostSignup).Methods("POST")
	router.HandleFunc("/login", PostLogin).Methods("POST")
	router.HandleFunc("/users", GetUsers).Methods("GET")
	router.HandleFunc("/users", PutUsers).Methods("PUT")
	err := http.ListenAndServe(":3000", router)

	//basic error handling
	if err != nil {
		log.Fatal(err)
	}

}
