package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type citiesResponse struct {
	Cities []string `json:"cities"` // Cities capitalised to export it, otherwise json.Marshal() will ignore it.
}

func cityHandler(res http.ResponseWriter, req *http.Request) {

	responseData := &citiesResponse{
		Cities: []string{
			"Amsterdam",
			"Bristol",
			"London",
			"Bangkok",
			"San Francisco",
		},
	}

	data, _ := json.MarshalIndent(responseData, "", "  ")
	res.Header().Set("Content-Type", "application/json; charset=utf-8")
	res.Write(data)
}

func defaultHandler(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/plain; charset=utf-8")

	res.Write([]byte("Hello World!"))
}

func main() {
	log.Println("Hello World!")
	log.Println("Listening on this host: http://localhost:5000")
	http.HandleFunc("/", defaultHandler)
	http.HandleFunc("/cities", cityHandler)
	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		log.Fatal("Unable to listen on port 5000 : ", err)
	}
}