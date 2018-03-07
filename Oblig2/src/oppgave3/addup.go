package main

import "fmt"

func main() {
	fmt.Println("Legg sammen to tall fra terminal og summer i en annen funksjon")
	fmt.Println(addUp(10, 10))
}

func addUp(tall, tall2 int) int{
	return tall + tall2
}
