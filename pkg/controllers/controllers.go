package controllers

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

var s = shortner.URLService{
	Elements:  "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ",
	COUNTER:   1000000000,
	LONGTOID:  map[string]int{},
	IDTOSMALL: map[int]string{}}

var URLRecords []URLRecord

var OutputFile string = "outputs/test.json"

type URLRecord struct {
	ID       int    `json:"id"`
	LongURL  string `json:"longurl"`
	ShortURL string `json:"shorturl"`
}

type uri struct {
	Url string `json:"url"`
}

func readfromFile() []URLRecord {

	// read our opened File as a byte array.
	jsonFile, err := os.Open(OutputFile)
	if err != nil {
		fmt.Println(err)
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'readRecords' which we defined above
	var readRecords []URLRecord
	json.Unmarshal(byteValue, &readRecords)
	jsonFile.Close()
	return readRecords
}

func CreateRecord(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//Parse the url from request body
	urlname := &uri{}
	utils.ParseBody(r, urlname)
	inputurl := (*urlname).Url

	//convert url to short url via LongToShort method
	short, id, existing := s.LongToShort(inputurl)
	recordEntry := URLRecord{id, inputurl, short}
	json.NewDecoder(r.Body).Decode(&recordEntry)

	//check if url does not exist in the file then
	//only append to records and return response
	if !existing {
		URLRecords = append(URLRecords, recordEntry)
		file, _ := json.MarshalIndent(URLRecords, "", " ")
		_ = ioutil.WriteFile(OutputFile, file, 0644)
	}
	json.NewEncoder(w).Encode(recordEntry)

}

func DeleteRecord(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//Fetch the delete id from path parameter
	vars := mux.Vars(r)
	id := vars["id"]
	readRecords := readfromFile()

	//Search the id and delete it from in memory records
	for i, rec := range URLRecords {
		if fmt.Sprint(rec.ID) == id {
			delete(s.IDTOSMALL, rec.ID)
			delete(s.LONGTOID, rec.LongURL)
			URLRecords = append(URLRecords[:i], URLRecords[i+1:]...)
		}
	}

	//Search the id and delete it from in file/DB records
	for index, record := range readRecords {
		if fmt.Sprint(record.ID) == id {
			nweRecords := append(readRecords[:index], readRecords[index+1:]...)

			//return full list of records with deleted id
			json.NewEncoder(w).Encode(nweRecords)
			file, _ := json.MarshalIndent(nweRecords, "", " ")
			_ = ioutil.WriteFile(OutputFile, file, 0644)
			return
		}
	}
	w.WriteHeader(404)
	fmt.Fprintf(w, "ERROR: id: %v not found", id)

}

func GetURLs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//Read records from file and return the response
	getRecords := readfromFile()
	fmt.Print(s, "\n\n")
	json.NewEncoder(w).Encode(getRecords)
}
