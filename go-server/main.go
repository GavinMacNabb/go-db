package main

import (
	"log"
	"net/http"
	"github.com/GavinMacNabb/go-db/router"
)

func main() {
	r := router.Router()
    log.Println("Listing for requests at http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}