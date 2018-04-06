package app

import (
	"testing"
)

func TestSeleciona(t *testing.T) {
	if err := LoadOrgaos(); err != nil {
		t.Fatal(err)
	}
	orgaos := seleciona("cdd falcao")
	if len(orgaos) != 1 {
		t.Fatal("deve haver apenas um órgão parecido com cdd falcão")
	}
	if orgaos[0].Mcu != "00027754" {
		t.Fatal("cdd falcao errado")
	}
}
