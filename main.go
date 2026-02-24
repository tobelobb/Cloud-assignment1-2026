package main

import (
	"log"
	"net/http"
	"countryinfo/handlers"
)

//sends to the different handlers based on the path
//starts server on port 8080 and sends starting message
func main() {
	http.HandleFunc("/countryinfo/v1/status/", handlers.StatusHandler)
	http.HandleFunc("/countryinfo/v1/info/", handlers.InfoHandler)
	http.HandleFunc("/countryinfo/v1/exchange/", handlers.ExchangeHandler)
	http.HandleFunc("/", handlers.RootHandler)
	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}