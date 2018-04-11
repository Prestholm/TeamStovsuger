package main

import (
	"net/http"
	"fmt"
)


func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, Client")
}

func main() {
	http.HandleFunc("/", sayHello)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}