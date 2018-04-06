package app

import (
	"net/http"
	"fmt"
	"service/modelo"
	"strings"
	"encoding/json"

	"github.com/arbovm/levenshtein"
	"io/ioutil"
	"sync"
)

type resultadoBuscaOrgao struct {
	Termo     string
	Resultado []modelo.Orgao
}

var orgaos []modelo.Orgao

type cacheOrgaos struct {
	sync.RWMutex
	mapa map[string][]modelo.Orgao
}

var cacheOrg = cacheOrgaos{mapa: make(map[string][]modelo.Orgao)}

func LoadOrgaos() error {
	var err error
	orgaos, err = modelo.Orgaos()
	return err
}

func Orgaos(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	os, ok := values["o"]
	if !ok {
		http.Error(w, "órgão não informado", 500)
		return
	}
	rs := make([]resultadoBuscaOrgao, 0)
	for _, o := range os {
		r := resultadoBuscaOrgao{
			Termo:     o,
			Resultado: seleciona(o),
		}
		rs = append(rs, r)
	}
	bytes, err := json.MarshalIndent(rs, "", " ")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Add("Content-type", "application/json; charset=utf8")
	fmt.Fprintln(w, string(bytes))
}

func OrgaosJson(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}
	in, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer r.Body.Close()
	var textos []string
	err = json.Unmarshal(in, &textos)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	rs := make([]resultadoBuscaOrgao, 0)
	for _, o := range textos {
		r := resultadoBuscaOrgao{
			Termo:     o,
			Resultado: seleciona(o),
		}
		rs = append(rs, r)
	}
	bytes, err := json.MarshalIndent(rs, "", " ")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Add("Content-type", "application/json; charset=utf8")
	fmt.Fprintln(w, string(bytes))
}

func seleciona(txt string) []modelo.Orgao {
	txt = strings.ToUpper(txt)
	cacheOrg.RLock()
	cacheRes, ok := cacheOrg.mapa[txt]
	cacheOrg.RUnlock()
	if ok {
		return cacheRes
	}
	menorDis := 100
	var proximos = make([]modelo.Orgao, 0)
	for _, o := range orgaos {
		d := levenshtein.Distance(txt, o.Nome)
		if d < menorDis {
			menorDis = d
			proximos = proximos[:0]
			proximos = append(proximos, o)
		} else if d == menorDis {
			proximos = append(proximos, o)
		}
	}
	cacheOrg.Lock()
	cacheOrg.mapa[txt] = proximos
	cacheOrg.Unlock()
	return proximos
}
