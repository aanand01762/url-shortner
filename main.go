package main

import (
	"log"
	"net/http"

	"github.com/aanand01762/url-shortner/pkg/routes"

	"github.com/gorilla/mux"
)

func main() {

	//Create a router and start the server at port 8080
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", r))
}
