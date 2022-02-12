package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fileserv := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileserv)

	fmt.Print("server started at port 8080\n")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
