package handlers

import (
    "countryinfo/models"
	"encoding/json"
	"io"
	"log"
	"net/http"
    "strings"
)

func InfoHandler(w http.ResponseWriter, r *http.Request) {
	// Get country name from url
    parts := strings.Split(r.URL.Path, "/")
    countryCode := parts[4]
    if len(parts) < 5 || len(countryCode) != 2 {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(map[string]string{"error": "country code required"})
        return
    }

	// Fetch from REST Countries API
	url := "http://129.241.150.113:8080/v3.1/alpha/" + countryCode
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Error fetching country data: %v", err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "failed to fetch data"})
		return
	}
	defer resp.Body.Close()

	// Check if response error
	if resp.StatusCode != http.StatusOK {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(resp.StatusCode)
		json.NewEncoder(w).Encode(map[string]string{"error": "country not found"})
		return
	}

	// Read and return API response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading response body: %v", err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "failed to read data"})
		return
	}

    // Parse the response
    var data []models.RestCountry
    if err := json.Unmarshal(body, &data); err != nil || len(data) == 0 {
        log.Printf("Error parsing country data: %v", err)
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode(map[string]string{"error": "failed to parse data"})
        return
    }

    c := data[0]

    result := models.CountryInfoResponse{
        Name:       c.Name.Common,
        Continents: c.Continents,
        Population: c.Population,
        Area:       c.Area,
        Languages:  c.Languages,
        Borders:    c.Borders,
        Flag:       c.Flags.PNG,
        Capital:    firstOrEmpty(c.Capital),
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(result)
}

//for getting the first capital only
func firstOrEmpty(arr []string) string {
    if len(arr) > 0 {
        return arr[0]
    }
    return ""
}
