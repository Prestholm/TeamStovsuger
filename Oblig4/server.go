package main

import (
	"net/http"
	"log"
	"html/template"
	"encoding/json"
	"time"
	"fmt"
	"strings"

	"io/ioutil"
)

const URL = "https://api.coinmarketcap.com/v1/ticker/"

var client =&http.Client{Timeout: 10*time.Second}

func main() {

	http.HandleFunc("/", startSide)
	http.HandleFunc("/info", basicHandler)
	http.HandleFunc("/list", liste)
	if err := http.ListenAndServe(":8080", nil); err !=nil{
		panic(err)
	}
	log.Println("Listening...")
}
func liste(w http.ResponseWriter, r *http.Request){
	coins := new(Coins)
	getData(URL, coins)
	tpl.ParseFiles("list.html"	)
	tpl.ExecuteTemplate(w, "list.html", coins)

}
func basicHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	n := r.FormValue("Name")
	d := struct {
		Name string
	}{
		Name: strings.ToLower(n),
	}
	tpl.ExecuteTemplate(w, "print.html", d)
	ID1 := r.PostFormValue("Name")
	res2, _ := http.Get(URL+ID1)
	body2, _ := ioutil.ReadAll(res2.Body)
	coin2 := Coins{}
	err2 := json.Unmarshal(body2, &coin2)
	if err2 !=nil {
		fmt.Println("error: ", err2)

	}
	err := r.ParseForm()
	if err != nil{
		panic(err)
	}
	
	tpl.ExecuteTemplate(w, "info.html", coin2)
}

var tpl *template.Template

func getData(url string, target interface{})error{
	//henter json-data fra url. Decoder denne dataen.
	res,err := client.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	return json.NewDecoder(res.Body).Decode(target)

}

func init() {
	tpl = template.Must(template.ParseGlob("Templates/*"))
}
func startSide(w http.ResponseWriter, r *http.Request){
	tpl.ParseFiles("form.html")
	tpl.ExecuteTemplate(w, "form.html", nil)
}

type Coins []struct {

	ID               string `json:"id"`
	Name             string `json:"name"`
	Symbol           string `json:"symbol"`
	Rank             string `json:"rank"`
	PriceUsd         string `json:"price_usd"`
	PriceBtc         string `json:"price_btc"`
	Two4HVolumeUsd   string `json:"24h_volume_usd"`
	MarketCapUsd     string `json:"market_cap_usd"`
	AvailableSupply  string `json:"available_supply"`
	TotalSupply      string `json:"total_supply"`
	MaxSupply        string `json:"max_supply"`
	PercentChange1H  string `json:"percent_change_1h"`
	PercentChange24H string `json:"percent_change_24h"`
	PercentChange7D  string `json:"percent_change_7d"`
	LastUpdated      string `json:"last_updated"`

}

