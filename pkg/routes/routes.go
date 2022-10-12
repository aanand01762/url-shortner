package routes

import (
	"github.com/aanand01762/url-shortner/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {

	//Define the routes for api calls
	router.HandleFunc("/records", controllers.CreateRecord).Methods("POST")
	router.HandleFunc("/records/{id}", controllers.DeleteRecord).Methods("DELETE")
	router.HandleFunc("/records", controllers.GetURLs)

}
