package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	io.WriteString(w, "This is my website\n")
}

func getHello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Got /hello request\n")
	io.WriteString(w, "Hello World\n")
}

func handleForm(w http.ResponseWriter, r *http.Request) {
	// this will look for application/x-www-form-urlencoded content type
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// once form has been parsed correctly. can obtain form fields via
	// request.FormValue("foo")

	id := r.FormValue("kidId")
	fmt.Printf("Kid Id is: %v", id)
}

func main() {
	http.HandleFunc("/", getRoot)
	http.HandleFunc("/hello", getHello)
	http.HandleFunc("/handleForm", handleForm)
	err := http.ListenAndServe(":8080", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
