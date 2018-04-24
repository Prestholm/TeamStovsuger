package main

import (
	"net/http"
	"fmt"
	"time"
	"log"
	"io/ioutil"
	"encoding/json"

	"html/template"
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
	tpl.ParseFiles("People.html")
	if err != nil{
		log.Print(err)
	}
	tpl.ExecuteTemplate(w, "People.html", people1)
	if err != nil{
		log.Fatal(err)
	}
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


	tpl.ParseFiles("Entries.html")
	if err != nil{
		log.Print(err)
	}
	tpl.ExecuteTemplate(w, "Entries.html", entries1)
	if err != nil{
		log.Fatal(err)
	}

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
	tpl.ParseFiles("fylke.html")
	if err != nil{
		log.Print(err)
	}
	tpl.ExecuteTemplate(w, "fylke.html", fylke1)
	if err != nil{
		log.Fatal(err)
	}
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
	tpl.ParseFiles("jon.html")
	if err != nil{
		log.Print(err)
	}
	tpl.ExecuteTemplate(w, "jon.html", jon1)
	if err != nil{
		log.Fatal(err)
	}
}
func side5(w http.ResponseWriter, r *http.Request){

	url := "https://hacker-news.firebaseio.com/v0/item/8863.json?print=pretty"
	spaceClient := http.Client{
		Timeout: time.Second * 2, // Maximum of 2 secs
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}


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


	tpl.ParseFiles("hotel.html")
	if err != nil{
		log.Print(err)
	}
	tpl.ExecuteTemplate(w, "hotel.html", hotel1)
	if err != nil{
		log.Fatal(err)
	}
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
type hotel struct {
	By          string `json:"by"`
	Descendants int    `json:"descendants"`
	ID          int    `json:"id"`
	Kids        []int  `json:"kids"`
	Score       int    `json:"score"`
	Time        int    `json:"time"`
	Title       string `json:"title"`
	Type        string `json:"type"`
	URL         string `json:"url"`
}
var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("template/*"))
}