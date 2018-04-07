package app

import (
	"net/http"
	"fmt"
	"service/modelo"
	"encoding/json"
	"sync"
	"strings"
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
	defaultCacheOrgaos.mutex.Lock()
	defaultCacheOrgaos.mapa[txt] = orgs
	defaultCacheOrgaos.mutex.Unlock()
	return orgs, nil
}
