package main
import (
	"os"
	"log"
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
)

func main() {
	addtofile()
	sumfromfile()
	readResult("resultat.txt")
}

func addtofile() {

	var nummer1 int
	var nummer2 int

	//I terminal så skriver brukeren inn 2 tall, nummer 1 & nummer 2
	fmt.Println("Skriv inn nummer: ")
	fmt.Scan(&nummer1)
	fmt.Println("Skriv inn nummer: ")
	fmt.Scan(&nummer2)

	//resultat.txt blir generert og tallene blir lagt inn
	file, err := os.Create("resultat.txt")
	if err != nil {
		log.Fatal("Error, kan ikke lage fil", err)
	}
	defer file.Close()

	f, err := os.OpenFile("resultat.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := fmt.Fprintf(f, "%d\n%d", nummer1, nummer2); err != nil {
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}

}

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}
	// nummer 1 og nummer 2 blir så lagt sammen og representert i en egen fil, eller i terminalen
func readResult(path string) {
	data, err := ioutil.ReadFile(path)
	checkErr(err)
	tempData := string(data)
	stringData := strings.Split(tempData, "\n")
	temp := stringData[len(stringData)-2]
	resultat, err := strconv.Atoi(temp)
	checkErr(err)

	fmt.Println("Resultat: ", resultat)
}
