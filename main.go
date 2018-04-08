package main

import (
	"service/app"
	"net/http"
	"log"
)

func main() {
	if err := app.LoadOrgaos(); err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/orgaos", app.Orgs)
	http.HandleFunc("/muitosorgaos", app.MuitosOrgs)
	http.ListenAndServe(":4000", nil)
}
