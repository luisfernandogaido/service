package main

import (
	"service/app"
	"net/http"
	"os"
	"log"
	"bufio"
	"encoding/json"
	"fmt"
)

func main() {
	server()
}

func server() {
	if err := app.LoadOrgaos(); err != nil {
		panic(err)
	}
	http.HandleFunc("/orgaos", app.Orgaos)
	http.HandleFunc("/orgaosjson", app.OrgaosJson)
	http.ListenAndServe(":4000", nil)
}

func teste() {
	f, err := os.Open("./docs/orgaos-buscar.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)
	textos := make([]string, 0)
	for scanner.Scan(){
		t := scanner.Text()
		textos = append(textos, t)
	}
	bytes, err:=json.MarshalIndent(textos, "", " ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(bytes))
}
