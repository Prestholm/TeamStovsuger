package main

import (
	"net/http"
	"fmt"
	"time"
	"log"
	"io/ioutil"
	"encoding/json"
)


func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, Client")
}

func main() {
	http.HandleFunc("/", sayHello)
	http.HandleFunc("/1", side1)

	//http.HandleFunc("/2", s2)
	//http.HandleFunc("/3", s3)
	//http.HandleFunc("/4", s4)
	//http.HandleFunc("/5", s5)


	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}


}

func side1(w http.ResponseWriter, r *http.Request){
	url := "http://api.open-notify.org/astros.json"

	spaceClient := http.Client{
		Timeout: time.Second * 2, // Maximum of 2 secs
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "spacecount-tutorial")

	res, getErr := spaceClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	people1 := people{}
	jsonErr := json.Unmarshal(body, &people1)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}


	fmt.Fprint(w, people1.Message)
}

type people struct {
	Message string `json:"message"`
	Number int `json:"number"`
}
