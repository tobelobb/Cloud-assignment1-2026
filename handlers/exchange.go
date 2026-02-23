package handlers

import (
	"countryinfo/models"
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

func ExchangeHandler(w http.ResponseWriter, r *http.Request) {

	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 5 || len(parts[4]) != 2 {
		http.Error(w, `{"error":"two-letter country code required"}`, http.StatusBadRequest)
		return
	}

	countryCode := strings.ToUpper(parts[4])

	// Fetch main country
	url := "http://129.241.150.113:8080/v3.1/alpha/" + countryCode
	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, `{"error":"failed to fetch country data"}`, http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		http.Error(w, `{"error":"country not found"}`, http.StatusNotFound)
		return
	}

	var countryData []models.RestCountry
	if err := json.NewDecoder(resp.Body).Decode(&countryData); err != nil || len(countryData) == 0 {
		http.Error(w, `{"error":"failed to parse country data"}`, http.StatusInternalServerError)
		return
	}

	country := countryData[0]

	// Extract base currency
	var baseCurrency string
	for code := range country.Currencies {
		baseCurrency = code
		break
	}

	// Fetch base currency table ONCE
	baseURL := "http://129.241.150.113:9090/currency/" + baseCurrency
	baseResp, err := http.Get(baseURL)
	if err != nil || baseResp.StatusCode != http.StatusOK {
		http.Error(w, `{"error":"failed to fetch base currency data"}`, http.StatusBadGateway)
		return
	}
	defer baseResp.Body.Close()

	baseBody, _ := io.ReadAll(baseResp.Body)

	// Parse base currency JSON safely
	var raw map[string]interface{}
	if err := json.Unmarshal(baseBody, &raw); err != nil {
		http.Error(w, `{"error":"failed to parse base currency data"}`, http.StatusInternalServerError)
		return
	}

	var baseMap map[string]float64

	// Case 1: {"base":"NOK","rates":{...}}
	if rates, ok := raw["rates"].(map[string]interface{}); ok {
		baseMap = make(map[string]float64)
		for k, v := range rates {
			if f, ok := v.(float64); ok {
				baseMap[k] = f
			}
		}
	}

	// Case 2: {"NOK":{"SEK":0.98,...}}
	if baseMap == nil {
		if inner, ok := raw[baseCurrency].(map[string]interface{}); ok {
			baseMap = make(map[string]float64)
			for k, v := range inner {
				if f, ok := v.(float64); ok {
					baseMap[k] = f
				}
			}
		}
	}

	if baseMap == nil {
		http.Error(w, `{"error":"failed to parse base currency data"}`, http.StatusInternalServerError)
		return
	}

	neighborExchanges := make(map[string]interface{})

	// Loop through borders
	for _, borderCode := range country.Borders {

		// Fetch neighbor country
		neighborURL := "http://129.241.150.113:8080/v3.1/alpha/" + borderCode
		neighborResp, err := http.Get(neighborURL)
		if err != nil || neighborResp.StatusCode != http.StatusOK {
			continue
		}
		defer neighborResp.Body.Close()

		var neighborData []models.RestCountry
		if err := json.NewDecoder(neighborResp.Body).Decode(&neighborData); err != nil || len(neighborData) == 0 {
			continue
		}

		neighbor := neighborData[0]

		// Extract neighbor currency
		var neighborCurrency string
		for code := range neighbor.Currencies {
			neighborCurrency = code
			break
		}
		if neighborCurrency == "" {
			continue
		}

		// Reverse lookup: base → neighbor
		rateFromBase, ok := baseMap[neighborCurrency]
		if !ok || rateFromBase == 0 {
			continue
		}

		// Invert it: neighbor → base
		rateToBase := 1 / rateFromBase

		neighborExchanges[borderCode] = map[string]interface{}{
			"currency":   neighborCurrency,
			"rateToBase": rateToBase,
		}
	}

	result := map[string]interface{}{
		"countryCode":  countryCode,
		"country":      country.Name.Common,
		"baseCurrency": baseCurrency,
		"exchanges":    neighborExchanges,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
