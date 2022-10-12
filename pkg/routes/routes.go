package routes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/aanand01762/url-shortner/pkg/shortner"
	"github.com/aanand01762/url-shortner/pkg/utils"
	"github.com/gorilla/mux"
)

type URLRecord struct {
	ID       int    `json:"id"`
	LongURL  string `json:"longurl"`
	ShortURL string `json:"shorturl"`
}

type uri struct {
	Url string `json:"url"`
}

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/records", getURLs).Methods("GET")
	router.HandleFunc("/record", createRecord).Methods("POST")
	router.HandleFunc("/records/{url}", deleteRecord).Methods("DELETE")
}
var URLRecords []URLRecord

func getURLs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(URLRecords)
}

func deleteRecord(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, record := range URLRecords {
		if record.LongURL == params["url"] {
			URLRecords = append(URLRecords[:index], URLRecords[index+1:]...)
			json.NewEncoder(w).Encode(URLRecords)
			break
		}
	}
}

var s shortner.URLService

func createRecord(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	urlname := &uri{}
	utils.ParseBody(r, urlname)
	inputurl := (*urlname).Url
	short, id, existing := s.LongToShort(inputurl)
	record := URLRecord{id, inputurl, short}
	fmt.Println(record)
	json.NewDecoder(r.Body).Decode(&record)
	if !existing {
		URLRecords = append(URLRecords, record)
	}
	json.NewEncoder(w).Encode(record)
	file, _ := json.MarshalIndent(URLRecords, "", " ")

	_ = ioutil.WriteFile("test.json", file, 0644)
}
