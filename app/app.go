package app

import "net/http"

func printJson(w http.ResponseWriter) error {
	return nil
}

func cors(w http.ResponseWriter) {
	allowedHeaders := "Content-type, Cache-Control"
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", allowedHeaders)
}
