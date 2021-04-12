package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type personal struct {
	Type     string `json:"type"`
	Value    string `json:"value"`
	Document struct {
		Ssn string `json:"ssn"`
	}
}

//main
func main() {
	text, err := InputData()
	if err != nil {
		return
	}
	fmt.Printf("this is return: %s\n", text)
}

//input data
func InputData() (string, error) {
	var inpval string
	fmt.Println("enter with type, gender or document: ")
	fmt.Scan(&inpval)

	text, err := fetcher(inpval)
	if err != nil {
		return "", err
	}
	return text, err
}

//finding data
func fetcher(p string) (string, error) {

	var retval string

	res, err := http.Get("https://www.vemprobananadaterra.com.br/json/example.json")

	templ := &personal{}
	if err != nil {
		return "", err
	}

	fmt.Printf("this is input: %s\n", p)

	if p == "type" {
		err = json.NewDecoder(res.Body).Decode(templ)
		retval = templ.Type
	} else if p == "gender" {
		err = json.NewDecoder(res.Body).Decode(templ)
		retval = templ.Value
	} else if p == "document" {
		err = json.NewDecoder(res.Body).Decode(templ)
		retval = templ.Document.Ssn
	} else {
		fmt.Printf("not found\n")
	}
	return retval, err
}
