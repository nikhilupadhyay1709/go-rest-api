package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	Con "go-rest-api/controllers"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Con.HomeLink)
	router.HandleFunc("/event", Con.CreateEvent).Methods("POST")
	router.HandleFunc("/events", Con.GetAllEvents).Methods("GET")
	router.HandleFunc("/events/{id}", Con.GetOneEvent).Methods("GET")
	router.HandleFunc("/events/{id}", Con.UpdateEvent).Methods("PATCH")
	router.HandleFunc("/events/{id}", Con.DeleteEvent).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))
}
