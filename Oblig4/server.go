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

const URL = "https://api.coinmarketcap.com/v1/ticker/" //Url til apien

/*Client varaible, setter 10 sec timeout på requests til endpointet. Foretrukket over å bruke standard http.Client,
som har en timeout på 0 sec. For de som er interessert så kan de lese mer her:
https://medium.com/@nate510/don-t-use-go-s-default-http-client-4804cb19f779
*/
var client =&http.Client{Timeout: 10*time.Second}

func main() {

	http.HandleFunc("/", startSide) // Startsiden
	http.HandleFunc("/info", basicHandler) // infosiden
	http.HandleFunc("/list", liste) // liste for alle cryptoene
	if err := http.ListenAndServe(":8080", nil); err !=nil{
		panic(err)
	}
	log.Println("Listening...")
}
func liste(w http.ResponseWriter, r *http.Request){
	coins = new(Coins)
	getData(URL, coins)
	tpl.ParseFiles("list.html"	)
	tpl.ExecuteTemplate(w, "list.html", coins) // kjører siden med coins som data, slik at den kan hente alt data URL har å gi på Coins{}

}
func basicHandler(w http.ResponseWriter, r *http.Request) {
//Sjekker om metoden ikke er post.
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
		//Setter opp for FORMs (videre implementert i Form.html
	n := r.FormValue("Name")
	//struct for forms
	d := struct {
		Name string
	}{
		Name: strings.ToLower(n),
	}
	tpl.ExecuteTemplate(w, "print.html", d)// lager overskriften med bildene takket være dataen som henter navn
	ID1 := r.PostFormValue("Name") // gjør navn om til string
	res2, _ := http.Get(URL+ID1) // legger til URL sammen med den coinen man vil ha data for
	body2, _ := ioutil.ReadAll(res2.Body)
	coin2 := Coins{}
	err2 := json.Unmarshal(body2, &coin2)
	if err2 !=nil {
		fmt.Println("error: ", err2)

	}
		//Gå gjennom det som ble fyllt ut i formsa.
	err := r.ParseForm()
	if err != nil{
		panic(err)
	}
	//execute indexen vår med begrenset data som bare inneholder det URLen skal ha
	tpl.ExecuteTemplate(w, "info.html", coin2)
}

var tpl *template.Template
func getData(){
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
//startsiden vår
func startSide(w http.ResponseWriter, r *http.Request){
	tpl.ParseFiles("form.html")
	tpl.ExecuteTemplate(w, "form.html", nil)
}
//Slice struct som inneholder alt av info om hver crypto.
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

