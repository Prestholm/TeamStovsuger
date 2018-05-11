

package main

import ("testing"
	"os"
	"log"
	"net/http"
	"net/http/httptest"

	"time"
)

func TestStartSide(t *testing.T) {
	req, err := http.NewRequest("GET", "https://api.coinmarketcap.com/v1/ticker/", nil)
	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(startSide)
	handler.ServeHTTP(recorder, req)
	status := recorder.Code
	if  status != http.StatusOK {
		t.Errorf("Unexpected return got %v wanted %v",
			status, http.StatusOK)
	}
}



func eksistensTest(filename string) bool {
	if _, err := os.Stat(filename); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
// template testene får man Fail pga de er ikke i samme mappe som server.go
func TestTemplate1(t *testing.T){
	r := eksistensTest("form.html")
	if r == false {
		t.Error("Finner ikke filen")
	}
}

func TestTemplate2(t *testing.T){
	r := eksistensTest("info.html")
	if r == false {
		t.Error("Finner ikke filen")
	}
}

func TestTemplate3(t *testing.T){
	r := eksistensTest("print.html")
	if r == false {
		t.Error("Finner ikke filen")
	}
}

func TestTemplate4(t *testing.T){
	r := eksistensTest("index1.html")
	if r == false {
		t.Error("Finner ikke filen")
	}
}
// tester om det er json
func TestGetJson(t *testing.T) {
	coins := new(Coins)
		server := getData("https://api.coinmarketcap.com/v1/ticker/", coins)
		time.Sleep(time.Second*1)
		if server != nil {
			t.Errorf("Finner ikke JSON")
		}

}

