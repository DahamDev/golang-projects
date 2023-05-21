package main

import (
	"net/http"
	"fmt"
	"io"
)

func main() {

	port := ":3333"

	http.HandleFunc("/", getRoot)
	http.HandleFunc("/form", handleForm)

	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println("Error occured. Shutting down server")
	} 
}


func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	io.WriteString(w, "This is my website!\n")
}

func handleForm(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /form request\n")
	io.WriteString(w, "This is the form \n")
}