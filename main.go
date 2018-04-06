package main

import (
	"net/http"
	"service/app"
)

func main() {
	if err := app.LoadOrgaos(); err != nil {
		panic(err)
	}
	http.HandleFunc("/orgaos", app.Orgaos)
	http.ListenAndServe(":4000", nil)
}
