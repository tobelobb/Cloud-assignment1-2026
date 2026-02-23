package handlers

import (
    "encoding/json"
    "log"
    "net/http"
    "time"
)

var startTime = time.Now()

func StatusHandler(w http.ResponseWriter, r *http.Request) {
    resp := map[string]interface{}{
        "restcountriesapi": 200,
        "currenciesapi":    200,
        "version":          "v1",
        "uptime":           int(time.Since(startTime).Seconds()),
    }

    w.Header().Set("Content-Type", "application/json")
    if err := json.NewEncoder(w).Encode(resp); err != nil {
        log.Printf("Error encoding response: %v", err)
        w.WriteHeader(http.StatusInternalServerError)
    }
}
