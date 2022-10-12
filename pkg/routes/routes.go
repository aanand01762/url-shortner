package routes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

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
	router.HandleFunc("/records", createRecord).Methods("POST")
	router.HandleFunc("/records/{id}", deleteRecord).Methods("DELETE")
	router.HandleFunc("/records", getURLs)

}
var URLRecords []URLRecord

func deleteRecord(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]

	readRecords := readfromFile()

	fmt.Println(id)
	for index, record := range readRecords {
		if fmt.Sprint(record.ID) == id {
			nweRecords := append(readRecords[:index], readRecords[index+1:]...)
			fmt.Println(nweRecords)
			json.NewEncoder(w).Encode(nweRecords)
			file, _ := json.MarshalIndent(nweRecords, "", " ")
			_ = ioutil.WriteFile("test.json", file, 0644)
			break
		}
	}
}

func getURLs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	getRecords := readfromFile()
	fmt.Println(getRecords)
	json.NewEncoder(w).Encode(getRecords)
}

var s shortner.URLService

func createRecord(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	urlname := &uri{}
	utils.ParseBody(r, urlname)
	inputurl := (*urlname).Url
	short, id, existing := s.LongToShort(inputurl)
	recordEntry := URLRecord{id, inputurl, short}
	fmt.Println(recordEntry)
	json.NewDecoder(r.Body).Decode(&recordEntry)
	if !existing {
		URLRecords = append(URLRecords, recordEntry)
	}
	json.NewEncoder(w).Encode(recordEntry)
	file, _ := json.MarshalIndent(URLRecords, "", " ")

	_ = ioutil.WriteFile("test.json", file, 0644)
}

func readfromFile() []URLRecord {
	jsonFile, err := os.Open("test.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we initialize our readRecords array
	var readRecords []URLRecord

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'readRecords' which we defined above
	json.Unmarshal(byteValue, &readRecords)

	// closing of our jsonFile so that we can parse it later on
	jsonFile.Close()

	return readRecords
}
