package main

import (
	"log"
	"net/http"

	"countryinfo/handlers"
)
//test
// API endpoints
const (
	restCountriesAPI = "http://129.241.150.113:8080/v3.1/"
	currencyAPI      = "http://129.241.150.113:9090/currency/"
)

func main() {
	// Register handlers - use trailing slashes for path matching
	http.HandleFunc("/countryinfo/v1/status/", handlers.StatusHandler)
	http.HandleFunc("/countryinfo/v1/info/", handlers.InfoHandler)
	http.HandleFunc("/countryinfo/v1/exchange/", handlers.ExchangeHandler)
	http.HandleFunc("/", handlers.RootHandler)
	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}