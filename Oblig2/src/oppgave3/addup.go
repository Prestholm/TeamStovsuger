package main

import (
	"flag"
	"fmt"

)

func main(){

	var numba int
	var numbb int

	fmt.Println("Skriv inn nummer: ")
	fmt.Scan(&numba)
	fmt.Println("Skriv inn nummer: ")
	fmt.Scan(&numbb)

	flag.Parse()

	c := make(chan int, 2)
	c <- numba
	c <- numbb

	b(c)
	
}

func b(c chan int){
	b,g := <- c, <- c
	c <- b+g
	fmt.Println("Resultat: ", b+g)

}
