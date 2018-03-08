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
	c := make(chan int, 2)
	c <- *numba
	c <- *numbb
	b(c)
	h := <-c
 fmt.Println(h)

}
func b(c chan int){
	b,g := <- c, <- c
	c <- b+g

}
