package main

import (
	"encoding/json"
	"log"
	"os"
)

type InputQuery struct {
	Query string `json:"sound"`
}

type OutputQuery struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func main() {
	var inputQuery InputQuery
	var outputQuery OutputQuery

	err := json.NewDecoder(os.Stdin).Decode(&inputQuery)
	if err != nil {
		log.Fatal(err)
	}

	if inputQuery.Query == "honk" {
		outputQuery.Code = "200"
		outputQuery.Message = "successful honk"
	} else {
		outputQuery.Code = "400"
		outputQuery.Message = "failed honk"
	}

	if err := json.NewEncoder(os.Stdout).Encode(outputQuery); err != nil {
		log.Fatal(err)
	}

}
