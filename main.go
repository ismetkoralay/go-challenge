package main

import (
	"goChallenge/api"
	"log"
	"net/http"
)

func main() {
	server := api.CreateServer()
	log.Fatal(http.ListenAndServe(":5000", server))
}
