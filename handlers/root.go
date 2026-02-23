package handlers

import (
	"fmt"
	"net/http"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w,
		"this is my assignment 1 cloud spring 2026. code is two letter code for a country no = norway \n\n"+
			"Available endpoints:\n"+
			"/countryinfo/v1/status/\n"+
			"/countryinfo/v1/info/{code}\n"+
			"/countryinfo/v1/exchange/{code}",
	)
}
