package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"encoding/json"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/soma", soma)
	http.ListenAndServe(":4000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "service!")
}

func soma(w http.ResponseWriter, r *http.Request) {
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer r.Body.Close()
	var numeros []int
	if err = json.Unmarshal(bytes, &numeros); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	var soma int
	for _, n := range numeros {
		soma += n
	}
	w.Header().Set("Content-type", "application/json; charset=utf8")
	ret := struct {
		Soma int
	}{
		soma,
	}
	js, err:= json.MarshalIndent(ret, "", " ")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	fmt.Fprintln(w, string(js))
}
