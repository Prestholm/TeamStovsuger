package main

import (
	"flag"
	"fmt"

)


func main(){
	numba := flag.Int("a", 100, "int")
	numbb := flag.Int("b", 20, "int")
	// satt standard nummere for a og b til 100 og 20
	//
	flag.Parse()
	// har man endret verdi i terminalen så settes det inn i Parse()
	c := make(chan int, 2)
	c <- *numba
	c <- *numbb
	// legger til tallene i channel c
	b(c)
	// kommuniserer med den andre funskjonen til å legge sammen de to tallene og henter det ut igjen via channels
	h := <-c
	//definerer h-verdien til noe som kan printes ut og legges frem.
 fmt.Println(h)

}
func b(c chan int){
	b,g := <- c, <- c
	c <- b+g

}
