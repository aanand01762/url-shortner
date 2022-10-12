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

var s shortner.URLService
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

	_ = ioutil.WriteFile(OutputFile, file, 0644)
}

func DeleteRecord(w http.ResponseWriter, r *http.Request) {

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
			_ = ioutil.WriteFile(OutputFile, file, 0644)
			break
		}
	}
}

func GetURLs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	getRecords := readfromFile()
	fmt.Println(getRecords)
	json.NewEncoder(w).Encode(getRecords)
}
