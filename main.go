package main

import (
	"service/app"
	"net/http"
)

func main() {
	http.HandleFunc("/orgaos", app.Orgs)
	http.ListenAndServe(":4000", nil)
}