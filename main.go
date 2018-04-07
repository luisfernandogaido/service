package main

import (
	"service/app"
	"net/http"
	"fmt"
)

func main() {
	if err := app.LoadOrgaos(); err != nil {
		panic(err)
	}
	http.HandleFunc("/orgaos", app.Orgaos)
	http.HandleFunc("/orgaosjson", app.OrgaosJson)
	http.HandleFunc("/teste", hello)
	http.ListenAndServe(":4000", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world")
}