package routes

import (
	"github.com/gorilla/mux"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/records", createRecord).Methods("POST")
	router.HandleFunc("/records/{id}", deleteRecord).Methods("DELETE")
	router.HandleFunc("/records", getURLs)

}
