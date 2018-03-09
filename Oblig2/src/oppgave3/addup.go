package main

import (
	"flag"
	"fmt"
	"syscall"
	"os"
	"os/signal"
)

func main(){
	//Ved bruk av SIGINT kan vi stoppe exectuablen, sånn at vi kan lese hva resultatet er
	sign :=make(chan os.Signal, 1)
	signal.Notify(sign, syscall.SIGINT)

	var nummer1 int
	var nummer2 int
	//Brukeren skriver inn 2 tall
	fmt.Println("Skriv inn et nummer: ")
	fmt.Scan(&nummer1)
	if nummer1 <= 0 {
		fmt.Println("Error, du må skrive inn et tall, eller et tall som er høyere enn 0")
	}
	fmt.Println("Skriv inn et nummer: ")
	fmt.Scan(&nummer2)
	if nummer2 <= 0 {
		fmt.Println("Error, du må skrive inn et tall, eller et tall som er høyere enn 0")
	}


	flag.Parse()

	c := make(chan int, 2)
	c <- nummer1
	c <- nummer2

	b(c)
	h := <-c
	if h <= 0 {
		fmt.Println("Error, du må skrive inn som er høyere enn 0")
	}
	fmt.Println("Resultatet er: ", h)
	run:=<-sign
	fmt.Print(run)
}
//Via flags blir disse 2 tallene lagt sammen og sendt tilbake til func main
func b(c chan int){
	b,g := <- c, <- c
	c <- b+g

}
