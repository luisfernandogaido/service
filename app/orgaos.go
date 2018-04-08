package app

import (
	"net/http"
	"fmt"
	"service/modelo"
	"encoding/json"
	"sync"
	"strings"

	"github.com/arbovm/levenshtein"
	"io/ioutil"
)

type resultadoBuscaOrgao struct {
	Termo     string
	Resultado []modelo.Orgao
}

type cacheOrgaos struct {
	mutex sync.RWMutex
	mapa  map[string][]modelo.Orgao
}

var defaultCacheOrgaos = cacheOrgaos{
	mapa: make(map[string][]modelo.Orgao),
}

var orgaos []modelo.Orgao

func LoadOrgaos() error {
	var err error
	orgaos, err = modelo.Orgaos()
	return err
}

func Orgs(w http.ResponseWriter, r *http.Request) {
	cors(w)
	values := r.URL.Query()
	os, ok := values["o"]
	if !ok {
		http.Error(w, "informe ao menos um órgão", 500)
		return
	}
	rs := make([]resultadoBuscaOrgao, 0)
	for _, o := range os {
		orgs, err := selecionaOrgaos(o)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		r := resultadoBuscaOrgao{
			Termo:     o,
			Resultado: orgs,
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
func MuitosOrgs(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Método não permitido.", http.StatusMethodNotAllowed)
		return
	}
	defer r.Body.Close()
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var os []string
	err = json.Unmarshal(bytes, &os)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	rs := make([]resultadoBuscaOrgao, 0)
	for _, o := range os {
		orgs, err := selecionaOrgaos(o)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		r := resultadoBuscaOrgao{
			Termo:     o,
			Resultado: orgs,
		}
		rs = append(rs, r)
	}
	bytes, err = json.MarshalIndent(rs, "", " ")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Add("Content-type", "application/json; charset=utf8")
	fmt.Fprintln(w, string(bytes))

}

func selecionaOrgaos(txt string) ([]modelo.Orgao, error) {
	txt = strings.ToLower(txt)
	defaultCacheOrgaos.mutex.RLock()
	orgs, ok := defaultCacheOrgaos.mapa[txt]
	defaultCacheOrgaos.mutex.RUnlock()
	if ok {
		return orgs, nil
	}
	orgs, err := modelo.OrgaosSeleciona(txt)
	if err != nil {
		return nil, err
	}
	if len(orgs) == 0 {
		orgs = selecionaOrgaosParecidos(txt)
	}
	defaultCacheOrgaos.mutex.Lock()
	defaultCacheOrgaos.mapa[txt] = orgs
	defaultCacheOrgaos.mutex.Unlock()
	return orgs, nil
}

func selecionaOrgaosParecidos(txt string) []modelo.Orgao {
	txt = strings.ToUpper(txt)
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
	return proximos
}
