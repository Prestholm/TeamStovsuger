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
	http.HandleFunc("/2", side2)
	http.HandleFunc("/3", side3)
	http.HandleFunc("/4", side4)
	http.HandleFunc("/5", side5)


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
	b, err := json.Marshal(people1)
	s:= string(b)

	fmt.Fprint(w, s)

}
func side2(w http.ResponseWriter, r *http.Request){

	url := "https://hotell.difi.no/api/json/stavanger/miljostasjoner"
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

	entries1 := entries{}
	jsonErr := json.Unmarshal(body, &entries1)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	b, err := json.Marshal(entries1)
	s:= string(b)

	fmt.Fprint(w, s)
}
func side3(w http.ResponseWriter, r *http.Request){

	url := "https://register.geonorge.no/api/subregister/sosi-kodelister/kartverket/fylkesnummer-alle.json"
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

	fylke1 := fylke{}
	jsonErr := json.Unmarshal(body, &fylke1)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	b, err := json.Marshal(fylke1)
	s:= string(b)

	fmt.Fprint(w, s)
}
func side4(w http.ResponseWriter, r *http.Request){

	url := "https://anapioficeandfire.com/api/characters/585"
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

	jon1 := JonUmber{}
	jsonErr := json.Unmarshal(body, &jon1)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	b, err := json.Marshal(jon1)
	s:= string(b)

	fmt.Fprint(w, s)
}
func side5(w http.ResponseWriter, r *http.Request){

	url := "https://hotell.difi.no/api/json/_all"
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

	hotel1 := hotel{}
	jsonErr := json.Unmarshal(body, &hotel1)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	b, err := json.Marshal(hotel1)
	s:= string(b)

	fmt.Fprint(w, s)
}


	type people struct {
	People []struct {
	Name  string `json:"name"`
	Craft string `json:"craft"`
	} `json:"people"`
	Number  int    `json:"number"`
	Message string `json:"message"`
	}
type entries struct {

	Entries []struct {
	Latitude    string `json:"latitude"`
	Navn        string `json:"navn"`
	Plast       string `json:"plast"`
	GlassMetall string `json:"glass_metall"`
	TekstilSko  string `json:"tekstil_sko,omitempty"`
	Longitude   string `json:"longitude"`
} `json:"entries"`
	Page  int `json:"page"`
	Pages int `json:"pages"`
	Posts int `json:"posts"`
}
type fylke struct {

	ID                 string `json:"id"`
	Label              string `json:"label"`
	Lang               string `json:"lang"`
	Contentsummary     string `json:"contentsummary"`
	Owner              string `json:"owner"`
	Manager            string `json:"manager"`
	ContainedItemClass string `json:"containedItemClass"`
	UUID               string `json:"uuid"`
	Containeditems     []struct {
	ID            string        `json:"id"`
	Label         string        `json:"label"`
	Lang          string        `json:"lang"`
	Itemclass     string        `json:"itemclass"`
	UUID          string        `json:"uuid"`
	Status        string        `json:"status"`
	Description   string        `json:"description"`
	Owner         string        `json:"owner"`
	VersionNumber int           `json:"versionNumber"`
	Versions      []interface{} `json:"versions"`
	LastUpdated   string        `json:"lastUpdated"`
	DateSubmitted string        `json:"dateSubmitted"`
	DateAccepted  string        `json:"dateAccepted"`
	Codevalue     string        `json:"codevalue,omitempty"`
	Narrower      []interface{} `json:"narrower"`
	AlertDate     string        `json:"AlertDate"`
	EffectiveDate string        `json:"EffectiveDate"`
	ValidFrom     string        `json:"ValidFrom,omitempty"`
	ValidTo       string        `json:"ValidTo,omitempty"`
} `json:"containeditems"`
	ContainedSubRegisters   []interface{} `json:"containedSubRegisters"`
	LastUpdated             string        `json:"lastUpdated"`
	SelectedDOKMunicipality string        `json:"SelectedDOKMunicipality"`
}
type JonUmber struct {
	URL         string        `json:"url"`
	Name        string        `json:"name"`
	Gender      string        `json:"gender"`
	Culture     string        `json:"culture"`
	Born        string        `json:"born"`
	Died        string        `json:"died"`
	Titles      []string      `json:"titles"`
	Aliases     []string      `json:"aliases"`
	Father      string        `json:"father"`
	Mother      string        `json:"mother"`
	Spouse      string        `json:"spouse"`
	Allegiances []string      `json:"allegiances"`
	Books       []string      `json:"books"`
	PovBooks    []interface{} `json:"povBooks"`
	TvSeries    []string      `json:"tvSeries"`
	PlayedBy    []string      `json:"playedBy"`
}
type hotel []struct {
	ShortName string `json:"shortName"`
	Name      string `json:"name"`
	Location  string `json:"location"`
	Updated   int    `json:"updated"`
	Dataset   bool   `json:"dataset"`
}
