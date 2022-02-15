package main

import (
	"fmt"
	"log"
	"net/http"
)

func formhandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "Error in parsing form err: %v", err)
		return
	}

	fmt.Fprintf(w, "POST request successfully executed\n")
	name := r.FormValue("name")
	add := r.FormValue("addr")

	fmt.Fprintf(w, "name : %s\n", name)
	fmt.Fprintf(w, "address: %s", add)
}

func main() {
	fileserve := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileserve)
	http.HandleFunc("/form", formhandler)

	fmt.Print("server started at 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
