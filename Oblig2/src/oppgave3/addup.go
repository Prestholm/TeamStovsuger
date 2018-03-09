package main

import (
	"flag"
	"fmt"

)

func main(){

	var numba int
	var numbb int
	//Brukeren skriver inn 2 tall
	fmt.Println("Skriv inn nummer: ")
	fmt.Scan(&numba)
	fmt.Println("Skriv inn nummer: ")
	fmt.Scan(&numbb)

	flag.Parse()

	c := make(chan int, 2)
	c <- numba
	c <- numbb

	b(c)
	h := <-c
	fmt.Println("Resultat: ", h)
}
	//Via flags blir disse 2 tallene lagt sammen
func b(c chan int){
	b,g := <- c, <- c
	c <- b+g

}
