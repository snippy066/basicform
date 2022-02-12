package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fileserve := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileserve)

	fmt.Print("server started at 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
