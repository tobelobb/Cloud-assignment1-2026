package handlers

import (
	"encoding/json"
	"net/http"
	"time"
)

var startTime = time.Now()

func StatusHandler(w http.ResponseWriter, r *http.Request) {
	var restCUurl = "http://129.241.150.113:8080/v3.1/all"
	var CurrUrl = "http://129.241.150.113:9090/currency/NOK"
	restStatus := check(restCUurl)
	currStatus := check(CurrUrl)

	resp := map[string]interface{}{
		"restcountriesapi": restStatus,
		"currenciesapi":    currStatus,
		"version":          "v1",
		"uptime":           int(time.Since(startTime).Seconds()),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

//if response returns response code, else 500 for server error.
func check(url string) int {
	client := http.Client{Timeout: 3 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		return 500
	}
	defer resp.Body.Close()
	return resp.StatusCode
}
